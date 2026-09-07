package main

import (
	"fmt"
	"strings"
	"time"

	"blockchain/core" // Imports the package you just created
)

func main() {
	var chain []core.Block

	fmt.Println("--- Mining Genesis Block... ---")
	genesisBlock := core.Block{
		Index:        0,
		Timestamp:    time.Now().UTC().String(),
		Data:         "Genesis Block",
		PreviousHash: strings.Repeat("0", 64),
		Nonce:        0,
	}
	core.MineBlock(&genesisBlock)
	chain = append(chain, genesisBlock)
	fmt.Printf("Mined! Hash: %s\nNonce (Attempts): %d\n\n", genesisBlock.Hash, genesisBlock.Nonce)

	fmt.Println("--- Mining Block #1... ---")
	secondBlock := core.GenerateBlock(chain[len(chain)-1], "Transaction: User A sends 50 Coins")
	chain = append(chain, secondBlock)
	fmt.Printf("Mined! Hash: %s\nNonce (Attempts): %d\n\n", secondBlock.Hash, secondBlock.Nonce)
}
