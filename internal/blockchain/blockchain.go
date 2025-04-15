package blockchain

import (
	"errors"
	"time"
)

type Blockchain struct {
	chain []Block
}

func NewBlockchain() *Blockchain {
	genesis := Block{
		Index:     0,
		Timestamp: time.Now(),
		Data:      "Genesis Block",
		PrevHash:  "",
	}
	genesis.Hash = calculateHash(genesis)

	return &Blockchain{
		chain: []Block{genesis},
	}
}

func (bc *Blockchain) AddBlock(data string) Block {
	prevBlock := bc.chain[len(bc.chain)-1]
	newBlock := Block{
		Index:     prevBlock.Index + 1,
		Timestamp: time.Now(),
		Data:      data,
		PrevHash:  prevBlock.Hash,
	}
	newBlock.Hash = calculateHash(newBlock)
	bc.chain = append(bc.chain, newBlock)
	return newBlock
}

func (bc *Blockchain) GetAllBlocks() []Block {
	return bc.chain
}

func (bc *Blockchain) IsValid() error {
	for i := 1; i < len(bc.chain); i++ {
		current := bc.chain[i]
		previous := bc.chain[i-1]

		if current.PrevHash != previous.Hash {
			return errors.New("blockchain is invalid: broken hash link")
		}
		if calculateHash(current) != current.Hash {
			return errors.New("blockchain is invalid: incorrect hash")
		}
	}
	return nil
}
