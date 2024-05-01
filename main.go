package main

import (
	"fmt"

	"github.com/martinyonatann/learn-blockchain-go/blockchains"
)

func main() {
	bc := blockchains.NewBlockchain()

	bc.AddBlock("Send 1 BTC to Martin")
	bc.AddBlock("Send 2 more BTC to Martin")

	for _, block := range bc.Blocks {
		fmt.Printf("Prev. hash : %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}
}
