package main

import (
	"os"

	"github.com/blockchain-city/foundational-blockchain/internal/api"
	"github.com/blockchain-city/foundational-blockchain/internal/blockchain"
	"github.com/blockchain-city/foundational-blockchain/pkg/logger"
)

func main() {
	logger.Info("Starting Blockchain Node...")

	var bc *blockchain.Blockchain
	if _, err := os.Stat("data/chain.json"); err == nil {
		loaded, err := blockchain.LoadFromDisk()
		if err != nil {
			logger.Error(err)
			return
		}
		bc = loaded
		logger.Info("Blockchain loaded from disk")
	} else {
		bc = blockchain.NewBlockchain()
		if err := bc.SaveToDisk(); err != nil {
			logger.Error(err)
			return
		}
		logger.Info("Genesis block created and saved")
	}

	server := api.NewServer(bc)
	server.Run(":8080")
}
