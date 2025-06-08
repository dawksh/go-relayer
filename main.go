package main

import (
	"context"
	"fmt"
	"os"

	"go-relayer/client"

	"github.com/ethereum/go-ethereum/common"
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

			for _, tx := range block.Transactions() {
				if tx.To().Cmp(common.HexToAddress("0x28172273CC1E0395F3473EC6eD062B6fdFb15940")) == 0 {
					fmt.Println("New transaction:", tx.Hash())
				}
			}
		}
	}
}
