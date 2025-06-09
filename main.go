package main

import (
	"context"
	"fmt"
	"os"

	"go-relayer/client"
	"go-relayer/utils"

	"github.com/ethereum/go-ethereum/core/types"
)

const contractABI = `[{"inputs":[],"name":"example","outputs":[],"stateMutability":"nonpayable","type":"function"}]`

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

	logger := utils.GetLogger()

	if err != nil {
		fmt.Println("Failed to connect to the Ethereum client:", err)
		os.Exit(1)
	}

	chainID, err := ethClient.ChainID(context.Background())
	if err != nil {
		fmt.Println("Failed to get chain ID:", err)
		os.Exit(1)
	}
	logger.Info("Connected to the Ethereum client with chain id:", chainID)

	header := make(chan *types.Header)
	sub, err := ethClient.SubscribeNewHead(context.Background(), header)

	if err != nil {
		fmt.Println("Failed to subscribe to new blocks:", err)
		os.Exit(1)
	}

	queue := make(chan *types.Block, 100)
	go func() {
		for block := range queue {
			go client.ProcessBlock(block)
		}
	}()

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

			queue <- block
		}
	}
}
