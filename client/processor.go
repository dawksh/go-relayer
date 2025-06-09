package client

import (
	"go-relayer/utils"

	"github.com/ethereum/go-ethereum/core/types"
)

func ProcessBlock(block *types.Block) {
	logger := utils.GetLogger()

	logger.Info("Processing block:", block.Number())

	for _, tx := range block.Transactions() {
		logger.Info("Transaction:", tx.Hash())
	}
}
