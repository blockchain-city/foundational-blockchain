package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

const Difficulty = 3

func mineBlock(block *Block) {
	nonce := 0
	for {
		block.Nonce = nonce
		block.Hash = calculateHash(*block)
		if isHashValid(block.Hash) {
			break
		}
		nonce++
	}
}

func isHashValid(hash string) bool {
	prefix := ""
	for i := 0; i < Difficulty; i++ {
		prefix += "0"
	}
	return hash[:Difficulty] == prefix
}

func calculateHash(block Block) string {
	record := fmt.Sprintf("%d%s%s%s%d",
		block.Index,
		block.Timestamp.Format("2006-01-02 15:04:05"),
		block.Data,
		block.PrevHash,
		block.Nonce,
	)
	h := sha256.New()
	h.Write([]byte(record))
	return hex.EncodeToString(h.Sum(nil))
}
