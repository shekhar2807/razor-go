//Package cmd provides all functions related to command line
package cmd

import (
	"razor/core"
	"razor/core/types"

	"github.com/ethereum/go-ethereum/common"
)

//This function allows the user to claim the block reward and returns the hash
func (*UtilsStruct) ClaimBlockReward(options types.TransactionOptions) (common.Hash, error) {
	epoch, err := razorUtils.GetEpoch(options.Client)
	if err != nil {
		log.Error("Error in getting epoch: ", err)
		return core.NilHash, err
	}
	log.Debug("ClaimBlockReward: Epoch: ", epoch)

	sortedProposedBlockIds, err := razorUtils.GetSortedProposedBlockIds(options.Client, epoch)
	if err != nil {
		log.Error("Error in getting sortedProposedBlockIds: ", err)
		return core.NilHash, err
	}
	log.Debug("ClaimBlockReward: Sorted proposed block Ids: ", sortedProposedBlockIds)

	if sortedProposedBlockIds == nil {
		log.Debug("No blocks proposed in this epoch")
		return core.NilHash, nil
	}

	stakerID, err := razorUtils.GetStakerId(options.Client, options.AccountAddress)
	if err != nil {
		log.Error("Error in getting stakerId: ", err)
		return core.NilHash, err
	}
	log.Debug("ClaimBlockReward: Staker Id: ", stakerID)

	selectedProposedBlock, err := razorUtils.GetProposedBlock(options.Client, epoch, sortedProposedBlockIds[0])
	if err != nil {
		log.Error("Error in getting selectedProposedBlock: ", err)
		return core.NilHash, err
	}
	log.Debug("ClaimBlockReward: Selected proposed block: ", selectedProposedBlock)

	if selectedProposedBlock.ProposerId == stakerID {
		log.Info("Claiming block reward...")
		txnOpts := razorUtils.GetTxnOpts(options)
		log.Debug("Executing ClaimBlockReward transaction...")
		txn, err := blockManagerUtils.ClaimBlockReward(options.Client, txnOpts)
		if err != nil {
			log.Error("Error in claiming block reward: ", err)
			return core.NilHash, err
		}
		log.Info("Txn Hash: ", transactionUtils.Hash(txn).Hex())
		return transactionUtils.Hash(txn), nil
	}

	log.Debug("Only selected block proposer can claim block reward")
	return core.NilHash, nil
}
