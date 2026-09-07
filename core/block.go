package core

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"strings"
	"time"
)

// Difficulty sets how many leading zeros the hash must have
const Difficulty = 3

// Block structure includes a Nonce for Proof-of-Work
type Block struct {
	Index        int
	Timestamp    string
	Data         string
	PreviousHash string
	Hash         string
	Nonce        int
}

func CalculateHash(b Block) string {
	record := strconv.Itoa(b.Index) + b.Timestamp + b.Data + b.PreviousHash + strconv.Itoa(b.Nonce)
	h := sha256.New()
	h.Write([]byte(record))
	return hex.EncodeToString(h.Sum(nil))
}

// MineBlock forces the computer to solve a cryptographic puzzle
func MineBlock(newBlock *Block) {
	target := strings.Repeat("0", Difficulty)
	for {
		newBlock.Hash = CalculateHash(*newBlock)
		if strings.HasPrefix(newBlock.Hash, target) {
			break
		}
		newBlock.Nonce++
	}
}

func GenerateBlock(oldBlock Block, data string) Block {
	var newBlock Block
	newBlock.Index = oldBlock.Index + 1
	newBlock.Timestamp = time.Now().UTC().String()
	newBlock.Data = data
	newBlock.PreviousHash = oldBlock.Hash
	newBlock.Nonce = 0

	MineBlock(&newBlock)
	return newBlock
}
