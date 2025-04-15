package blockchain

import (
	"encoding/json"
	"os"
)

const dataFile = "data/chain.json"

func (bc *Blockchain) SaveToDisk() error {
	file, err := os.Create(dataFile)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	return encoder.Encode(bc.chain)
}

func LoadFromDisk() (*Blockchain, error) {
	file, err := os.Open(dataFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var chain []Block
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&chain); err != nil {
		return nil, err
	}
	return &Blockchain{chain: chain}, nil
}
