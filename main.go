package main

import (
	"blockchain/block"
	"blockchain/block/validators"
	"blockchain/prime_number"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	
	blockchain := block.InitBlockchain()

	go func() {
		<-sigs
		for _, b := range blockchain.GetBlocks() {
			fmt.Println(b)
		}
		if !validators.Validate(blockchain) {
			fmt.Println("BC not valid")
		}
		os.Exit(0)
	}()


	ch := make(chan int)
	go prime_number.Generate(ch)

	for {
		prime := <-ch
		//fmt.Println("Add in BC:", prime)
		blockchain.AddNewBlock([]byte(strconv.Itoa(prime)))
		ch1 := make(chan int)
		go prime_number.Filter(ch, ch1, prime)
		ch = ch1	
	}

	

}
