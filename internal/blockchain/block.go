package blockchain

import "time"

type Block struct {
	Index     int       `json:"index"`
	Timestamp time.Time `json:"timestamp"`
	Data      string    `json:"data"`
	PrevHash  string    `json:"prev_hash"`
	Hash      string    `json:"hash"`
}
