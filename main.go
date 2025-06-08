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
	fmt.Println(` ________   ________          ________   _______    ___        ________       ___    ___  _______    ________     
|\   ____\ |\   __  \        |\   __  \ |\  ___ \  |\  \      |\   __  \     |\  \  /  /||\  ___ \  |\   __  \    
\ \  \___| \ \  \|\  \       \ \  \|\  \\ \   __/| \ \  \     \ \  \|\  \    \ \  \/  / /\ \   __/| \ \  \|\  \   
 \ \  \  ___\ \  \\\  \       \ \   _  _\\ \  \_|/__\ \  \     \ \   __  \    \ \    / /  \ \  \_|/__\ \   _  _\  
  \ \  \|\  \\ \  \\\  \       \ \  \\  \|\ \  \_|\ \\ \  \____ \ \  \ \  \    \/  /  /    \ \  \_|\ \\ \  \\  \| 
   \ \_______\\ \_______\       \ \__\\ _\ \ \_______\\ \_______\\ \__\ \__\ __/  / /       \ \_______\\ \__\\ _\ 
    \|_______| \|_______|        \|__|\|__| \|_______| \|_______| \|__|\|__||\___/ /         \|_______| \|__|\|__|
                                                                            \|___|/                               
                                                                                                                  
                                                                                                                  `)

	ethClient, err := client.GetEthClient()

	if err != nil {
		fmt.Println("Failed to connect to the Ethereum client:", err)
		os.Exit(1)
	}

	chainID, err := ethClient.ChainID(context.Background())
	if err != nil {
		fmt.Println("Failed to get chain ID:", err)
		os.Exit(1)
	}
	fmt.Println("Connected to the Ethereum client with chain id:", chainID)

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
				if to := tx.To(); to != nil && to.Cmp(common.HexToAddress("0x28172273CC1E0395F3473EC6eD062B6fdFb15940")) == 0 {
					fmt.Println("New transaction:", common.Bytes2Hex(tx.Data()))
				}
			}
		}
	}
}
