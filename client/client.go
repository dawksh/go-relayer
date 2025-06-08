package client

import (
	"os"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
}

func GetEthClient() (*ethclient.Client, error) {
	return ethclient.Dial(os.Getenv("BASE_RPC"))
}
