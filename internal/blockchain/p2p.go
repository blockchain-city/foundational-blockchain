package blockchain

import (
	"encoding/json"
	"io"
	"net/http"
)

func (bc *Blockchain) SyncFromNode(nodeURL string) error {
	resp, err := http.Get(nodeURL + "/chain")
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var remoteChain []Block
	if err := json.Unmarshal(body, &remoteChain); err != nil {
		return err
	}

	if len(remoteChain) > len(bc.chain) {
		bc.chain = remoteChain
	}
	return nil
}
