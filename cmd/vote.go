//Package cmd provides all functions related to command line
package cmd

import (
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"os"
	"os/signal"
	"path"
	"razor/accounts"
	"razor/core"
	"razor/core/types"
	"razor/logger"
	"razor/pkg/bindings"
	"razor/utils"
	"time"

	"github.com/ethereum/go-ethereum/crypto"

	"github.com/spf13/pflag"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	solsha3 "github.com/miguelmota/go-solidity-sha3"
	"github.com/spf13/cobra"
)

var voteCmd = &cobra.Command{
	Use:   "vote",
	Short: "Start monitoring contract, commit, reveal, propose and dispute automatically",
	Long: `vote command allows you to participate in the voting of assets and earn rewards.

Example:
  ./razor vote --address 0x5a0b54d5dc17e0aadc383d2db43b0a0d3e029c4c`,
	Run: initializeVote,
}

//This function initialises the ExecuteVote function
func initializeVote(cmd *cobra.Command, args []string) {
	cmdUtils.ExecuteVote(cmd.Flags())
}

//This function sets the flag appropriately and executes the Vote function
func (*UtilsStruct) ExecuteVote(flagSet *pflag.FlagSet) {
	config, err := cmdUtils.GetConfigData()
	utils.CheckError("Error in getting config: ", err)
	log.Debugf("ExecuteVote: Config: %+v", config)

	client := razorUtils.ConnectToClient(config.Provider)

	address, err := flagSetUtils.GetStringAddress(flagSet)
	utils.CheckError("Error in getting address: ", err)
	log.Debug("ExecuteVote: Address: ", address)

	logger.SetLoggerParameters(client, address)
	log.Debug("Checking to assign log file...")
	razorUtils.AssignLogFile(flagSet)

	log.Debug("Getting password...")
	password := razorUtils.AssignPassword(flagSet)

	isRogue, err := flagSetUtils.GetBoolRogue(flagSet)
	utils.CheckError("Error in getting rogue status: ", err)
	log.Debug("ExecuteVote: IsRogue: ", isRogue)

	rogueMode, err := flagSetUtils.GetStringSliceRogueMode(flagSet)
	utils.CheckError("Error in getting rogue modes: ", err)
	log.Debug("ExecuteVote: RogueMode: ", rogueMode)

	backupNodeActionsToIgnore, err := flagSetUtils.GetStringSliceBackupNode(flagSet)
	utils.CheckError("Error in getting backupNode actions to ignore: ", err)
	log.Debug("ExecuteVote: Backup node actions to ignore: ", backupNodeActionsToIgnore)

	rogueData := types.Rogue{
		IsRogue:   isRogue,
		RogueMode: rogueMode,
	}

	if rogueData.IsRogue {
		log.Warn("YOU ARE RUNNING VOTE IN ROGUE MODE, THIS CAN INCUR PENALTIES!")
	}

	account := types.Account{Address: address, Password: password}

	cmdUtils.HandleExit()
	log.Debugf("Calling Vote() with arguments rogueData = %+v, account address = %s, backup node actions to ignore = %s", rogueData, account.Address, backupNodeActionsToIgnore)
	if err := cmdUtils.Vote(context.Background(), config, client, rogueData, account, backupNodeActionsToIgnore); err != nil {
		log.Errorf("%v\n", err)
		osUtils.Exit(1)
	}
}

//This function handles the exit and listens for CTRL+C
func (*UtilsStruct) HandleExit() {
	// listen for CTRL+C
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	defer func() {
		signal.Stop(signalChan)
		cancel()
	}()

	go func() {
		select {
		case <-signalChan:
			log.Warn("If you don't unstake and withdraw your coins, you may get inactivity penalty!")
			log.Info("Press CTRL+C again to terminate.")
		case <-ctx.Done():
		}
		<-signalChan // second signal, hard exit
		os.Exit(2)
	}()
}

//This function handles all the states of voting
func (*UtilsStruct) Vote(ctx context.Context, config types.Configurations, client *ethclient.Client, rogueData types.Rogue, account types.Account, backupNodeActionsToIgnore []string) error {
	header, err := utils.UtilsInterface.GetLatestBlockWithRetry(client)
	utils.CheckError("Error in getting block: ", err)
	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			log.Debugf("Vote: Header value: %d", header.Number)
			latestHeader, err := utils.UtilsInterface.GetLatestBlockWithRetry(client)
			if err != nil {
				log.Error("Error in fetching block: ", err)
				continue
			}
			log.Debugf("Vote: Latest header value: %d", latestHeader.Number)
			if latestHeader.Number.Cmp(header.Number) != 0 {
				header = latestHeader
				cmdUtils.HandleBlock(client, account, latestHeader.Number, config, rogueData, backupNodeActionsToIgnore)
			}
			time.Sleep(time.Second * time.Duration(core.BlockNumberInterval))
		}
	}
}

var (
	globalCommitDataStruct types.CommitFileData
	lastVerification       uint32
	blockConfirmed         uint32
	disputeData            types.DisputeFileData
)

//This function handles the block
func (*UtilsStruct) HandleBlock(client *ethclient.Client, account types.Account, blockNumber *big.Int, config types.Configurations, rogueData types.Rogue, backupNodeActionsToIgnore []string) {
	state, err := razorUtils.GetDelayedState(client, config.BufferPercent)
	if err != nil {
		log.Error("Error in getting state: ", err)
		return
	}
	epoch, err := razorUtils.GetEpoch(client)
	if err != nil {
		log.Error("Error in getting epoch: ", err)
		return
	}

	stakerId, err := razorUtils.GetStakerId(client, account.Address)
	if err != nil {
		log.Error("Error in getting staker id: ", err)
		return
	}
	if stakerId == 0 {
		log.Error("Staker doesn't exist")
		return
	}
	staker, err := razorUtils.GetStaker(client, stakerId)
	if err != nil {
		log.Error(err)
		return
	}
	stakedAmount := staker.Stake

	ethBalance, err := utils.UtilsInterface.BalanceAtWithRetry(client, common.HexToAddress(account.Address))
	if err != nil {
		log.Errorf("Error in fetching balance of the account: %s\n%v", account.Address, err)
		return
	}
	actualStake, err := razorUtils.ConvertWeiToEth(stakedAmount)
	if err != nil {
		log.Error("Error in converting stakedAmount from wei denomination: ", err)
		return
	}
	actualBalance, err := razorUtils.ConvertWeiToEth(ethBalance)
	if err != nil {
		log.Error("Error in converting ethBalance from wei denomination: ", err)
		return
	}

	sRZRBalance, err := razorUtils.GetStakerSRZRBalance(client, staker)
	if err != nil {
		log.Error("Error in getting sRZR balance for staker: ", err)
		return
	}

	var sRZRInEth *big.Float
	if sRZRBalance.Cmp(big.NewInt(0)) == 0 {
		sRZRInEth = big.NewFloat(0)
	} else {
		sRZRInEth, err = razorUtils.ConvertWeiToEth(sRZRBalance)
		if err != nil {
			log.Error(err)
			return
		}
	}

	log.Infof("State: %s Staker ID: %d Stake: %f sRZR Balance: %f Eth Balance: %f", utils.UtilsInterface.GetStateName(state), stakerId, actualStake, sRZRInEth, actualBalance)

	if staker.IsSlashed {
		log.Error("Staker is slashed.... cannot continue to vote!")
		osUtils.Exit(0)
	}

	switch state {
	case 0:
		log.Debugf("Starting commit...")
		err := cmdUtils.InitiateCommit(client, config, account, epoch, stakerId, rogueData)
		if err != nil {
			log.Error(err)
			break
		}
	case 1:
		log.Debugf("Starting reveal...")
		err := cmdUtils.InitiateReveal(client, config, account, epoch, staker, rogueData)
		if err != nil {
			log.Error(err)
			break
		}
	case 2:
		log.Debugf("Starting propose...")
		err := cmdUtils.InitiatePropose(client, config, account, epoch, staker, blockNumber, rogueData)
		if err != nil {
			log.Error(err)
			break
		}
	case 3:
		log.Debugf("Last verification: %d", lastVerification)
		if lastVerification >= epoch {
			log.Debugf("Last verification (%d) is greater or equal to current epoch (%d)", lastVerification, epoch)
			log.Debugf("Won't dispute now")
			break
		}

		err := cmdUtils.HandleDispute(client, config, account, epoch, blockNumber, rogueData, backupNodeActionsToIgnore)
		if err != nil {
			log.Error(err)
			break
		}

		lastVerification = epoch

		if utilsInterface.IsFlagPassed("autoClaimBounty") {
			log.Debugf("Automatically claiming bounty")
			err = cmdUtils.HandleClaimBounty(client, config, account)
			if err != nil {
				log.Error(err)
				break
			}
		}

	case 4:
		log.Debugf("Last verification: %d", lastVerification)
		log.Debugf("Block confirmed: %d", blockConfirmed)
		if lastVerification == epoch && blockConfirmed < epoch {
			txn, err := cmdUtils.ClaimBlockReward(types.TransactionOptions{
				Client:          client,
				Password:        account.Password,
				AccountAddress:  account.Address,
				ChainId:         core.ChainId,
				Config:          config,
				ContractAddress: core.BlockManagerAddress,
				MethodName:      "claimBlockReward",
				ABI:             bindings.BlockManagerABI,
			})

			if err != nil {
				log.Error("ClaimBlockReward error: ", err)
				break
			}
			if txn != core.NilHash {
				waitForBlockCompletionErr := razorUtils.WaitForBlockCompletion(client, txn.Hex())
				if waitForBlockCompletionErr != nil {
					log.Error("Error in WaitForBlockCompletion for claimBlockReward: ", err)
					break
				}
				blockConfirmed = epoch
			}
		}
	case -1:
		if config.WaitTime > 5 {
			timeUtils.Sleep(5 * time.Second)
			return
		}
	}
	razorUtils.WaitTillNextNSecs(config.WaitTime)
	fmt.Println()
}

//This function initiates the commit
func (*UtilsStruct) InitiateCommit(client *ethclient.Client, config types.Configurations, account types.Account, epoch uint32, stakerId uint32, rogueData types.Rogue) error {
	staker, err := razorUtils.GetStaker(client, stakerId)
	if err != nil {
		log.Error(err)
		return err
	}
	log.Debug("InitiateCommit: Staker:", staker)
	stakedAmount := staker.Stake
	minStakeAmount, err := utils.UtilsInterface.GetMinStakeAmount(client)
	if err != nil {
		log.Error("Error in getting minimum stake amount: ", err)
		return err
	}
	log.Debug("InitiateCommit: Minimum stake amount: ", minStakeAmount)
	if stakedAmount.Cmp(minStakeAmount) < 0 {
		log.Error("Stake is below minimum required. Kindly add stake to continue voting.")
		return nil
	}

	lastCommit, err := razorUtils.GetEpochLastCommitted(client, stakerId)
	if err != nil {
		return errors.New("Error in fetching last commit: " + err.Error())
	}
	log.Debug("InitiateCommit: Epoch last committed: ", lastCommit)

	if lastCommit >= epoch {
		log.Debugf("Cannot commit in epoch %d because last committed epoch is %d", epoch, lastCommit)
		return nil
	}
	razorPath, err := razorUtils.GetDefaultPath()
	if err != nil {
		return err
	}
	log.Debugf("InitiateCommit: .razor directory path: %s", razorPath)
	keystorePath := path.Join(razorPath, "keystore_files")
	log.Debugf("InitiateCommit: Keystore file path: %s", keystorePath)
	log.Debugf("InitiateCommit: Calling CalculateSecret() with arguments epoch = %d, keystorePath = %s, chainId = %s", epoch, keystorePath, core.ChainId)
	_, secret, err := cmdUtils.CalculateSecret(account, epoch, keystorePath, core.ChainId)
	if err != nil {
		return err
	}
	log.Debug("InitiateCommit: Secret: ", secret)

	log.Debugf("Getting Salt for current epoch %d...", epoch)
	salt, err := cmdUtils.GetSalt(client, epoch)
	if err != nil {
		return err
	}
	log.Debug("InitiateCommit: Salt: ", salt)

	seed := solsha3.SoliditySHA3([]string{"bytes32", "bytes32"}, []interface{}{"0x" + hex.EncodeToString(salt[:]), "0x" + hex.EncodeToString(secret)})

	log.Debugf("InitiateCommit: Calling HandleCommitState with arguments epoch = %d, seed = %v, rogueData = %+v", epoch, seed, rogueData)
	commitData, err := cmdUtils.HandleCommitState(client, epoch, seed, rogueData)
	if err != nil {
		return errors.New("Error in getting active assets: " + err.Error())
	}
	log.Debug("InitiateCommit: Commit Data: ", commitData)

	log.Debug("InitiateCommit: Calling CreateMerkle() with argument Leaves = ", commitData.Leaves)
	merkleTree := utils.MerkleInterface.CreateMerkle(commitData.Leaves)
	log.Debug("InitiateCommit: Merkle Tree: ", merkleTree)
	commitTxn, err := cmdUtils.Commit(client, config, account, epoch, seed, utils.MerkleInterface.GetMerkleRoot(merkleTree))
	if err != nil {
		return errors.New("Error in committing data: " + err.Error())
	}
	log.Debug("InitiateCommit: Commit Transaction Hash: ", commitTxn)
	if commitTxn != core.NilHash {
		waitForBlockCompletionErr := razorUtils.WaitForBlockCompletion(client, commitTxn.String())
		if waitForBlockCompletionErr != nil {
			log.Error("Error in WaitForBlockCompletion for commit: ", err)
			return errors.New("error in sending commit transaction")
		}
		log.Debug("Updating GlobalCommitDataStruct with latest commitData and epoch...")
		updateGlobalCommitDataStruct(commitData, epoch)
		log.Debugf("InitiateCommit: Global commit data struct: %+v", globalCommitDataStruct)
	}

	log.Debug("Saving committed data for recovery...")
	fileName, err := razorUtils.GetCommitDataFileName(account.Address)
	if err != nil {
		return errors.New("Error in getting file name to save committed data: " + err.Error())
	}
	log.Debug("InitiateCommit: Commit data file path: ", fileName)

	err = razorUtils.SaveDataToCommitJsonFile(fileName, epoch, commitData)
	if err != nil {
		return errors.New("Error in saving data to file" + fileName + ": " + err.Error())
	}
	log.Debug("Data saved!")
	return nil
}

//This function initiates the reveal
func (*UtilsStruct) InitiateReveal(client *ethclient.Client, config types.Configurations, account types.Account, epoch uint32, staker bindings.StructsStaker, rogueData types.Rogue) error {
	stakedAmount := staker.Stake
	log.Debug("InitiateReveal: Staked Amount: ", stakedAmount)
	minStakeAmount, err := utils.UtilsInterface.GetMinStakeAmount(client)
	if err != nil {
		log.Error("Error in getting minimum stake amount: ", err)
		return err
	}
	log.Debug("InitiateReveal: Minimum Stake Amount: ", minStakeAmount)
	if stakedAmount.Cmp(minStakeAmount) < 0 {
		log.Error("Stake is below minimum required. Kindly add stake to continue voting.")
		return nil
	}
	lastReveal, err := razorUtils.GetEpochLastRevealed(client, staker.Id)
	if err != nil {
		return errors.New("Error in fetching last reveal: " + err.Error())
	}
	log.Debug("InitiateReveal: Last reveal was at epoch ", lastReveal)

	if lastReveal >= epoch {
		log.Debugf("Since last reveal was at epoch: %d, won't reveal again in epoch: %d", lastReveal, epoch)
		return nil
	}

	log.Debugf("InitiateReveal: Calling HandleRevealState with arguments staker = %+v, epoch = %d", staker, epoch)
	if err := cmdUtils.HandleRevealState(client, staker, epoch); err != nil {
		log.Error(err)
		return err
	}

	nilCommitData := globalCommitDataStruct.AssignedCollections == nil && globalCommitDataStruct.SeqAllottedCollections == nil && globalCommitDataStruct.Leaves == nil

	if nilCommitData {
		log.Debug("InitiateReveal: Global commit data is nil, getting the commit data from file...")
		fileName, err := razorUtils.GetCommitDataFileName(account.Address)
		if err != nil {
			log.Error("Error in getting file name to save committed data: ", err)
			return err
		}
		log.Debug("InitiateReveal: Commit data file path: ", fileName)
		committedDataFromFile, err := razorUtils.ReadFromCommitJsonFile(fileName)
		if err != nil {
			log.Errorf("Error in getting committed data from file %s: %v", fileName, err)
			return err
		}
		log.Debug("InitiateReveal: Committed data from file: ", committedDataFromFile)
		if committedDataFromFile.Epoch != epoch {
			log.Errorf("File %s doesn't contain latest committed data", fileName)
			return errors.New("commit data file doesn't contain latest committed data")
		}
		log.Debug("Updating global commit data struct...")
		updateGlobalCommitDataStruct(types.CommitData{
			Leaves:                 committedDataFromFile.Leaves,
			SeqAllottedCollections: committedDataFromFile.SeqAllottedCollections,
			AssignedCollections:    committedDataFromFile.AssignedCollections,
		}, epoch)
		log.Debugf("InitiateReveal: Global Commit data struct: %+v", globalCommitDataStruct)
	}
	if rogueData.IsRogue && utils.Contains(rogueData.RogueMode, "reveal") {
		log.Warn("YOU ARE REVEALING VALUES IN ROGUE MODE, THIS CAN INCUR PENALTIES!")
		var rogueCommittedData []*big.Int
		for i := 0; i < len(globalCommitDataStruct.Leaves); i++ {
			rogueCommittedData = append(rogueCommittedData, razorUtils.GetRogueRandomValue(10000000))
		}
		globalCommitDataStruct.Leaves = rogueCommittedData
		log.Debugf("InitiateReveal: Global Commit data struct in rogue mode: %+v", globalCommitDataStruct)
	}

	if globalCommitDataStruct.Epoch == epoch {
		razorPath, err := razorUtils.GetDefaultPath()
		if err != nil {
			return err
		}
		log.Debug("InitiateReveal: .razor directory path: ", razorPath)
		keystorePath := path.Join(razorPath, "keystore_files")
		log.Debug("InitiateReveal: Keystore file path: ", keystorePath)

		log.Debugf("InitiateReveal: Calling CalculateSecret() with argument epoch = %d, keystorePath = %s, chainId = %s", epoch, keystorePath, core.ChainId)
		signature, _, err := cmdUtils.CalculateSecret(account, epoch, keystorePath, core.ChainId)
		if err != nil {
			return err
		}
		log.Debug("InitiateReveal: Signature: ", signature)
		log.Debug("InitiateReveal: Assigned Collections: ", globalCommitDataStruct.AssignedCollections)
		log.Debug("InitiateReveal: SeqAllottedCollections: ", globalCommitDataStruct.SeqAllottedCollections)
		log.Debug("InitiateReveal: Leaves: ", globalCommitDataStruct.Leaves)

		commitDataToSend := types.CommitData{
			Leaves:                 globalCommitDataStruct.Leaves,
			AssignedCollections:    globalCommitDataStruct.AssignedCollections,
			SeqAllottedCollections: globalCommitDataStruct.SeqAllottedCollections,
		}
		log.Debugf("InitiateReveal: Calling Reveal() with arguments epoch = %d, commitDataToSend = %+v, signature = %v", epoch, commitDataToSend, signature)
		revealTxn, err := cmdUtils.Reveal(client, config, account, epoch, commitDataToSend, signature)
		if err != nil {
			return errors.New("Reveal error: " + err.Error())
		}
		if revealTxn != core.NilHash {
			waitForBlockCompletionErr := razorUtils.WaitForBlockCompletion(client, revealTxn.String())
			if waitForBlockCompletionErr != nil {
				log.Error("Error in WaitForBlockCompletionErr for reveal: ", err)
				return err
			}
		}
	} else {
		log.Error("The commit data is outdated, does not match with the latest epoch")
		return errors.New("outdated commit data")
	}
	return nil
}

//This function initiates the propose
func (*UtilsStruct) InitiatePropose(client *ethclient.Client, config types.Configurations, account types.Account, epoch uint32, staker bindings.StructsStaker, blockNumber *big.Int, rogueData types.Rogue) error {
	stakedAmount := staker.Stake
	log.Debug("InitiatePropose: Staked Amount: ", stakedAmount)
	minStakeAmount, err := utils.UtilsInterface.GetMinStakeAmount(client)
	if err != nil {
		log.Error("Error in getting minimum stake amount: ", err)
		return err
	}
	log.Debug("InitiatePropose: Minimum Stake Amount: ", minStakeAmount)
	if stakedAmount.Cmp(minStakeAmount) < 0 {
		log.Error("Stake is below minimum required. Kindly add stake to continue voting.")
		return nil
	}
	lastProposal, err := razorUtils.GetEpochLastProposed(client, staker.Id)
	if err != nil {
		return errors.New("Error in fetching last proposal: " + err.Error())
	}
	log.Debug("InitiatePropose: Last propose was in epoch ", lastProposal)
	if lastProposal >= epoch {
		log.Debugf("Since last propose was at epoch: %d, won't propose again in epoch: %d", epoch, lastProposal)
		return nil
	}
	lastReveal, err := razorUtils.GetEpochLastRevealed(client, staker.Id)
	if err != nil {
		return errors.New("Error in fetching last reveal: " + err.Error())
	}
	log.Debug("InitiatePropose: Last reveal was in epoch ", lastReveal)
	if lastReveal < epoch {
		log.Debugf("Cannot propose in epoch %d because last reveal was in epoch %d", epoch, lastReveal)
		return nil
	}

	log.Debugf("InitiatePropose: Calling Propose() with arguments staker = %+v, epoch = %d, blockNumber = %s, rogueData = %+v", staker, epoch, blockNumber, rogueData)
	err = cmdUtils.Propose(client, config, account, staker, epoch, blockNumber, rogueData)
	if err != nil {
		return errors.New("Propose error: " + err.Error())
	}
	return nil
}

//This function calculates the secret
func (*UtilsStruct) CalculateSecret(account types.Account, epoch uint32, keystorePath string, chainId *big.Int) ([]byte, []byte, error) {
	if chainId == nil {
		return nil, nil, errors.New("chainId is nil")
	}
	hash := solsha3.SoliditySHA3([]string{"address", "uint32", "uint256", "string"}, []interface{}{common.HexToAddress(account.Address), epoch, chainId, "razororacle"})
	log.Debug("CalculateSecret: Hash: ", hash)
	ethHash := utils.SignHash(hash)
	log.Debug("Hash generated for secret")
	log.Debug("CalculateSecret: Ethereum signed hash: ", ethHash)
	signedData, err := accounts.AccountUtilsInterface.SignData(ethHash, account, keystorePath)
	if err != nil {
		return nil, nil, errors.New("Error in signing the data: " + err.Error())
	}
	log.Debug("CalculateSecret: SignedData: ", signedData)
	log.Debugf("Checking whether recovered address from Hash: %v and Signed data: %v is same as given address...", hash, signedData)
	recoveredAddress, err := utils.EcRecover(hash, signedData)
	if err != nil {
		return nil, nil, errors.New("Error in verifying: " + err.Error())
	}
	log.Debug("CalculateSecret: Recovered Address: ", recoveredAddress)
	if recoveredAddress != common.HexToAddress(account.Address) {
		return nil, nil, errors.New("invalid verification")
	}
	log.Debug("Address verified, generating secret....")
	if signedData[64] == 0 || signedData[64] == 1 {
		signedData[64] += 27
	}

	secret := crypto.Keccak256(signedData)
	log.Debug("Secret generated.")
	return signedData, secret, nil
}

func updateGlobalCommitDataStruct(commitData types.CommitData, epoch uint32) types.CommitFileData {
	globalCommitDataStruct.Leaves = commitData.Leaves
	globalCommitDataStruct.AssignedCollections = commitData.AssignedCollections
	globalCommitDataStruct.SeqAllottedCollections = commitData.SeqAllottedCollections
	globalCommitDataStruct.Epoch = epoch
	return globalCommitDataStruct
}

func init() {
	rootCmd.AddCommand(voteCmd)

	var (
		Address         string
		Rogue           bool
		RogueMode       []string
		Password        string
		AutoClaimBounty bool
		BackupNode      []string
	)

	voteCmd.Flags().StringVarP(&Address, "address", "a", "", "address of the staker")
	voteCmd.Flags().BoolVarP(&Rogue, "rogue", "r", false, "enable rogue mode to report wrong values")
	voteCmd.Flags().StringSliceVarP(&RogueMode, "rogueMode", "", []string{}, "type of rogue mode")
	voteCmd.Flags().StringVarP(&Password, "password", "", "", "password path of the staker to protect the keystore")
	voteCmd.Flags().BoolVarP(&AutoClaimBounty, "autoClaimBounty", "", false, "auto claim bounty")
	voteCmd.Flags().StringSliceVarP(&BackupNode, "backupNode", "", []string{}, "actions that backup node will ignore")

	addrErr := voteCmd.MarkFlagRequired("address")
	utils.CheckError("Address error: ", addrErr)
}
