package main

import (
	"context"
	"fmt"
	"os"

	"go-relayer/client"

	"github.com/ethereum/go-ethereum/core/types"
)

func main() {
	fmt.Println(`GO RELAYER`)

	ethClient, err := client.GetEthClient()

	if err != nil {
		fmt.Println("Failed to connect to the Ethereum client:", err)
		os.Exit(1)
	}

	header := make(chan *types.Header)
	sub, err := ethClient.SubscribeNewHead(context.Background(), header)

	if err != nil {
		fmt.Println("Failed to subscribe to new blocks:", err)
		os.Exit(1)
	}

	for {
		select {
		case err := <-sub.Err():
			fmt.Println("Subscription error:", err)

		case header := <-header:
			block, err := ethClient.BlockByHash(context.Background(), header.Hash())
			if err != nil {
				fmt.Println("Failed to get block by hash:", err)
				continue
			}

			fmt.Println("New block:", block.Number())
		}
	}
}
