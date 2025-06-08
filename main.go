package main

import (
	"context"
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	fmt.Println("go-relayer")
	ethClient, err := ethclient.Dial("https://boldest-damp-pond.base-mainnet.quiknode.pro/")

	if err != nil {
		fmt.Println("Failed to connect to the Ethereum client:", err)
		os.Exit(1)
	}

	block, err := ethClient.BlockNumber(context.Background())
	if err != nil {
		fmt.Println("Failed to get the latest block:", err)
		os.Exit(1)
	}

	fmt.Println("Latest block:", block)
	os.Exit(0)
}
