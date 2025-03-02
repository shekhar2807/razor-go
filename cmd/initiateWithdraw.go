//Package cmd provides all functions related to command line
package cmd

import (
	"errors"
	"math/big"
	"razor/core"
	"razor/core/types"
	"razor/logger"
	"razor/pkg/bindings"
	"razor/utils"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var initiateWithdrawCmd = &cobra.Command{
	Use:   "initiateWithdraw",
	Short: "initiateWithdraw for your razors once you've unstaked",
	Long: `initiateWithdraw command can be used once the user has unstaked their token and the withdraw period is upon them.

Example:
  ./razor initiateWithdraw --address 0x5a0b54d5dc17e0aadc383d2db43b0a0d3e029c4c --stakerId 1
`,
	Run: func(cmd *cobra.Command, args []string) {
		cmdUtils.ExecuteInitiateWithdraw(cmd.Flags())
	},
}

//This function sets the flags appropriately and executes the InitiateWithdraw function
func (*UtilsStruct) ExecuteInitiateWithdraw(flagSet *pflag.FlagSet) {
	config, err := cmdUtils.GetConfigData()
	utils.CheckError("Error in getting config: ", err)
	log.Debugf("ExecuteInitiateWithdraw: Config: %+v: ", config)

	client := razorUtils.ConnectToClient(config.Provider)

	address, err := flagSetUtils.GetStringAddress(flagSet)
	utils.CheckError("Error in getting address: ", err)
	log.Debug("ExecuteInitiateWithdraw: Address: ", address)

	logger.SetLoggerParameters(client, address)
	log.Debug("Checking to assign log file...")
	razorUtils.AssignLogFile(flagSet)

	log.Debug("Getting password...")
	password := razorUtils.AssignPassword(flagSet)

	razorUtils.CheckEthBalanceIsZero(client, address)

	stakerId, err := razorUtils.AssignStakerId(flagSet, client, address)
	utils.CheckError("Error in fetching stakerId:  ", err)
	log.Debug("ExecuteInitiateWithdraw: Staker Id: ", stakerId)

	log.Debugf("ExecuteInitiateWithdraw: Calling HandleUnstakeLock() with arguments account address: %s, stakerId: %d", address, stakerId)
	txn, err := cmdUtils.HandleUnstakeLock(client, types.Account{
		Address:  address,
		Password: password,
	}, config, stakerId)

	utils.CheckError("InitiateWithdraw error: ", err)
	if txn != core.NilHash {
		err := razorUtils.WaitForBlockCompletion(client, txn.String())
		utils.CheckError("Error in WaitForBlockCompletion for initiateWithdraw: ", err)
	}
}

//This function handles the unstake lock
func (*UtilsStruct) HandleUnstakeLock(client *ethclient.Client, account types.Account, configurations types.Configurations, stakerId uint32) (common.Hash, error) {
	_, err := cmdUtils.WaitForAppropriateState(client, "initiateWithdraw", 0, 1, 4)
	if err != nil {
		log.Error("Error in fetching epoch: ", err)
		return core.NilHash, err
	}

	unstakeLock, err := razorUtils.GetLock(client, account.Address, stakerId, 0)
	if err != nil {
		log.Error("Error in fetching unstakeLock")
		return core.NilHash, err
	}
	log.Debugf("HandleUnstakeLock: Unstake lock: %+v", unstakeLock)

	if unstakeLock.UnlockAfter.Cmp(big.NewInt(0)) == 0 {
		log.Error("Unstake command not called before initiating withdrawal!")
		return core.NilHash, errors.New("unstake Razors before withdrawing")
	}

	withdrawInitiationPeriod, err := razorUtils.GetWithdrawInitiationPeriod(client)
	if err != nil {
		log.Error("Error in fetching withdraw release period")
		return core.NilHash, err
	}
	log.Debug("HandleUnstakeLock: Withdraw initiation period: ", withdrawInitiationPeriod)

	withdrawBefore := big.NewInt(0).Add(unstakeLock.UnlockAfter, big.NewInt(int64(withdrawInitiationPeriod)))
	log.Debug("HandleUnstakeLock: Withdraw before epoch: ", withdrawBefore)
	epoch, err := razorUtils.GetEpoch(client)
	if err != nil {
		log.Error("Error in fetching epoch")
		return core.NilHash, err
	}
	log.Debug("HandleUnstakeLock: Current epoch: ", epoch)

	if big.NewInt(int64(epoch)).Cmp(withdrawBefore) > 0 {
		log.Info("Withdraw initiation period has passed. Cannot withdraw now, please reset the unstakeLock!")
		return core.NilHash, nil
	}

	waitFor := big.NewInt(0).Sub(unstakeLock.UnlockAfter, big.NewInt(int64(epoch)))
	if waitFor.Cmp(big.NewInt(0)) > 0 {
		timeRemaining := int64(uint32(waitFor.Int64())) * core.EpochLength
		if waitFor.Cmp(big.NewInt(1)) == 0 {
			log.Infof("Withdrawal period not reached. Cannot withdraw now, please wait for %d epoch! (approximately %s)", waitFor, razorUtils.SecondsToReadableTime(int(timeRemaining)))
		} else {
			log.Infof("Withdrawal period not reached. Cannot withdraw now, please wait for %d epochs! (approximately %s)", waitFor, razorUtils.SecondsToReadableTime(int(timeRemaining)))
		}
		return core.NilHash, nil
	}

	txnArgs := types.TransactionOptions{
		Client:          client,
		Password:        account.Password,
		AccountAddress:  account.Address,
		ChainId:         core.ChainId,
		Config:          configurations,
		ContractAddress: core.StakeManagerAddress,
		MethodName:      "initiateWithdraw",
		ABI:             bindings.StakeManagerABI,
		Parameters:      []interface{}{stakerId},
	}
	txnOpts := razorUtils.GetTxnOpts(txnArgs)

	if big.NewInt(int64(epoch)).Cmp(unstakeLock.UnlockAfter) >= 0 && big.NewInt(int64(epoch)).Cmp(withdrawBefore) <= 0 {
		log.Debug("Calling InitiateWithdraw() with arguments stakerId: ", stakerId)
		return cmdUtils.InitiateWithdraw(client, txnOpts, stakerId)
	}
	return core.NilHash, errors.New("unstakeLock period not over yet! Please try after some time")
}

//This function initiate withdraw for your razors once you've unstaked
func (*UtilsStruct) InitiateWithdraw(client *ethclient.Client, txnOpts *bind.TransactOpts, stakerId uint32) (common.Hash, error) {
	log.Info("Initiating withdrawal of funds...")
	log.Debug("Executing InitiateWithdraw transaction for stakerId = ", stakerId)
	txn, err := stakeManagerUtils.InitiateWithdraw(client, txnOpts, stakerId)
	if err != nil {
		log.Error("Error in initiating withdrawal of funds")
		return core.NilHash, err
	}

	log.Info("Txn Hash: ", transactionUtils.Hash(txn))

	return transactionUtils.Hash(txn), nil
}

func init() {
	rootCmd.AddCommand(initiateWithdrawCmd)

	var (
		Address  string
		Password string
		StakerId uint32
	)

	initiateWithdrawCmd.Flags().StringVarP(&Address, "address", "a", "", "address of the user")
	initiateWithdrawCmd.Flags().StringVarP(&Password, "password", "", "", "password path of user to protect the keystore")
	initiateWithdrawCmd.Flags().Uint32VarP(&StakerId, "stakerId", "", 0, "password path of user to protect the keystore")

	addrErr := initiateWithdrawCmd.MarkFlagRequired("address")
	utils.CheckError("Address error: ", addrErr)
}
