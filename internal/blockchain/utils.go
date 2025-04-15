package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

func calculateHash(block Block) string {
	record := fmt.Sprintf("%d%s%s%s",
		block.Index,
		block.Timestamp.Format(time.RFC3339),
		block.Data,
		block.PrevHash,
	)

	h := sha256.New()
	h.Write([]byte(record))
	return hex.EncodeToString(h.Sum(nil))
}
