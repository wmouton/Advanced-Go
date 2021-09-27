package main

import (
	"fmt"
	"github.com/Advanced-Go/Day-3/5-Developing-a-Blockchain-SDK/blockchain"
)

func main() {
	bc := blockchain.NewBlockchain(blockchain.NewGenesisBlock())
	fmt.Println(bc.GetCurrentBlock().Hash)
	fmt.Println(blockchain.GetTransactionHash(*bc.GetCurrentBlock().Transaction))
	bc.AddBlock(*blockchain.NewTransaction([]byte{4, 5, 6}))
	fmt.Println(bc.GetCurrentBlock().PreviousHash)
	fmt.Println(bc.GetCurrentBlock().Hash)
}
