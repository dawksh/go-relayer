package client

import (
	"github.com/ethereum/go-ethereum/ethclient"
)

func GetEthClient() (*ethclient.Client, error) {
	return ethclient.Dial("https://boldest-damp-pond.base-mainnet.quiknode.pro/")
}
