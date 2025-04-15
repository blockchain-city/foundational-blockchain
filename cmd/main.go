package main

import (
	"fmt"

	"github.com/blockchain-city/foundational-blockchain/internal/blockchain"
	"github.com/blockchain-city/foundational-blockchain/pkg/logger"
)

func main() {
	logger.Info("Initializing Blockchain...")
	bc := blockchain.NewBlockchain()

	logger.Info("Adding blocks...")
	bc.AddBlock("اولین داده")
	bc.AddBlock("دومین داده")
	bc.AddBlock("سومین داده")

	if err := bc.IsValid(); err != nil {
		logger.Error(err)
	} else {
		logger.Info("Blockchain is valid.")
	}

	for _, block := range bc.GetAllBlocks() {
		fmt.Printf("Index: %d\nTime: %s\nData: %s\nHash: %s\nPrevHash: %s\n\n",
			block.Index, block.Timestamp, block.Data, block.Hash, block.PrevHash)
	}
}
