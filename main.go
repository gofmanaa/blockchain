package main

import (
	"blockchain/block"
	"fmt"
)

func main() {
	blockchain := block.InitBlockchain()
	
	blockchain.AddNewBlock([]byte("test 0"))
	blockchain.AddNewBlock([]byte("test 1"))
	blockchain.AddNewBlock([]byte("test 2"))
	blockchain.AddNewBlock([]byte("test 3"))

	for _, b := range blockchain.GetBlocks() {
		fmt.Println(b)
	}

}
