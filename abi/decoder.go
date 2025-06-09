package abi

import (
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

func DecodeTransactionData(contractABI string, txData string) (string, error) {
	parsedABI, err := abi.JSON(strings.NewReader(contractABI))
	if err != nil {
		return "", fmt.Errorf("failed to parse ABI: %v", err)
	}

	data, err := hex.DecodeString(strings.TrimPrefix(txData, "0x"))
	if err != nil {
		return "", fmt.Errorf("failed to decode tx data: %v", err)
	}

	if len(data) < 4 {
		return "", fmt.Errorf("tx data too short")
	}

	methodID := data[:4]
	method, err := parsedABI.MethodById(methodID)
	if err != nil {
		return "", fmt.Errorf("failed to get method: %v", err)
	}

	args, err := method.Inputs.Unpack(data[4:])
	if err != nil {
		return "", fmt.Errorf("failed to unpack args: %v", err)
	}

	return fmt.Sprintf("Method: %s\nArgs: %v", method.Name, args), nil
}
