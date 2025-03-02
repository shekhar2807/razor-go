//Package cmd provides all functions related to command line
package cmd

import (
	"razor/core"
	"razor/core/types"
	"razor/logger"
	"razor/pkg/bindings"
	"razor/utils"

	"github.com/spf13/pflag"

	"github.com/ethereum/go-ethereum/common"

	"github.com/spf13/cobra"
)

var stakeCmd = &cobra.Command{
	Use:   "addStake",
	Short: "Stake some razors",
	Long: `addStake allows user to stake razors in the razor network

Example:
  ./razor addStake --address 0x5a0b54d5dc17e0aadc383d2db43b0a0d3e029c4c --value 1000 --logFile addStake`,
	Run: initialiseStake,
}

//This function initialises the ExecuteStake function
func initialiseStake(cmd *cobra.Command, args []string) {
	cmdUtils.ExecuteStake(cmd.Flags())
}

//This function sets the flags appropriately and executes the StakeCoins function
func (*UtilsStruct) ExecuteStake(flagSet *pflag.FlagSet) {
	config, err := cmdUtils.GetConfigData()
	utils.CheckError("Error in getting config: ", err)
	log.Debugf("ExecuteStake: config: %+v", config)

	client := razorUtils.ConnectToClient(config.Provider)

	address, err := flagSetUtils.GetStringAddress(flagSet)
	utils.CheckError("Error in getting address: ", err)
	log.Debug("ExecuteStake: Address: ", address)

	logger.SetLoggerParameters(client, address)
	log.Debug("Checking to assign log file...")
	razorUtils.AssignLogFile(flagSet)

	log.Debug("Getting password...")
	password := razorUtils.AssignPassword(flagSet)

	balance, err := razorUtils.FetchBalance(client, address)
	utils.CheckError("Error in fetching razor balance for account: "+address, err)
	log.Debug("Getting amount in wei...")
	valueInWei, err := cmdUtils.AssignAmountInWei(flagSet)
	utils.CheckError("Error in getting amount: ", err)
	log.Debug("ExecuteStake: Amount in wei: ", valueInWei)

	log.Debug("Checking for sufficient balance...")
	razorUtils.CheckAmountAndBalance(valueInWei, balance)

	log.Debug("Checking whether sFuel balance is not 0...")
	razorUtils.CheckEthBalanceIsZero(client, address)

	minSafeRazor, err := utils.UtilsInterface.GetMinSafeRazor(client)
	utils.CheckError("Error in getting minimum safe razor amount: ", err)
	log.Debug("ExecuteStake: Minimum razor that you can stake for first time: ", minSafeRazor)

	stakerId, err := razorUtils.GetStakerId(client, address)
	utils.CheckError("Error in getting stakerId: ", err)
	log.Debug("ExecuteStake: Staker Id: ", stakerId)

	if valueInWei.Cmp(minSafeRazor) < 0 && stakerId == 0 {
		log.Fatal("The amount of razors entered is below min safe value.")
	}

	txnArgs := types.TransactionOptions{
		Client:         client,
		AccountAddress: address,
		Password:       password,
		Amount:         valueInWei,
		ChainId:        core.ChainId,
		Config:         config,
	}

	log.Debug("ExecuteStake: Calling Approve() for amount: ", txnArgs.Amount)
	approveTxnHash, err := cmdUtils.Approve(txnArgs)
	utils.CheckError("Approve error: ", err)

	if approveTxnHash != core.NilHash {
		err = razorUtils.WaitForBlockCompletion(txnArgs.Client, approveTxnHash.String())
		utils.CheckError("Error in WaitForBlockCompletion for approve: ", err)
	}

	log.Debug("ExecuteStake: Calling StakeCoins() for amount: ", txnArgs.Amount)
	stakeTxnHash, err := cmdUtils.StakeCoins(txnArgs)
	utils.CheckError("Stake error: ", err)

	err = razorUtils.WaitForBlockCompletion(txnArgs.Client, stakeTxnHash.String())
	utils.CheckError("Error in WaitForBlockCompletion for stake: ", err)
}

//This function allows the user to stake razors in the razor network and returns the hash
func (*UtilsStruct) StakeCoins(txnArgs types.TransactionOptions) (common.Hash, error) {
	epoch, err := razorUtils.GetEpoch(txnArgs.Client)
	if err != nil {
		return common.Hash{0x00}, err
	}
	log.Debug("StakeCoins: Epoch: ", epoch)

	txnArgs.ContractAddress = core.StakeManagerAddress
	txnArgs.MethodName = "stake"
	txnArgs.Parameters = []interface{}{epoch, txnArgs.Amount}
	txnArgs.ABI = bindings.StakeManagerABI
	txnOpts := razorUtils.GetTxnOpts(txnArgs)
	log.Debugf("Executing Stake transaction with epoch = %d, amount = %d", epoch, txnArgs.Amount)
	tx, err := stakeManagerUtils.Stake(txnArgs.Client, txnOpts, epoch, txnArgs.Amount)
	if err != nil {
		return common.Hash{0x00}, err
	}
	log.Info("Txn Hash: ", transactionUtils.Hash(tx).Hex())
	return transactionUtils.Hash(tx), nil
}

func init() {
	rootCmd.AddCommand(stakeCmd)
	var (
		Amount   string
		Address  string
		Password string
		WeiRazor bool
	)

	stakeCmd.Flags().StringVarP(&Amount, "value", "v", "0", "amount of Razors to stake")
	stakeCmd.Flags().StringVarP(&Address, "address", "a", "", "address of the staker")
	stakeCmd.Flags().StringVarP(&Password, "password", "", "", "password path of staker to protect the keystore")
	stakeCmd.Flags().BoolVarP(&WeiRazor, "weiRazor", "", false, "value can be passed in wei")

	amountErr := stakeCmd.MarkFlagRequired("value")
	utils.CheckError("Value error: ", amountErr)
	addrErr := stakeCmd.MarkFlagRequired("address")
	utils.CheckError("Address error: ", addrErr)

}
