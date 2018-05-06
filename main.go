package main

import (
	"fmt"
	"simplebc/blockchain"
	"strconv"
)

func main() {
	bc := blockchain.NewBlockChain()
	bc.AddBlock("Send 1 BTC to Ivan")
	bc.AddBlock("Send 2 more BTC to Ivan")
	for _, block := range bc.GetBlocks() {
		fmt.Printf("Previous Hash:%x\n", block.PrevHash)
		fmt.Printf("Data:%s\n", block.Data)
		fmt.Printf("Hash:%x\n", block.Hash)
		pow := blockchain.NewProofOfWork(block)
		fmt.Printf("PoW:%s", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}
}
