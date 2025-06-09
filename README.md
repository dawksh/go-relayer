# go relayer

A performant Ethereum relayer written in Go. It connects to an Ethereum-compatible node, subscribes to new blocks, and decodes transactions to a specific contract.

## Prerequisites

-   Go 1.23+
-   Access to an Ethereum-compatible RPC endpoint (e.g., Base, Infura, Alchemy)

## Setup

1. Clone the repository.
2. Create a `.env` file in the project root with:

    BASE_RPC=YOUR_ETHEREUM_RPC_URL

3. Install dependencies:

    go mod tidy

## Run

    go run main.go

The relayer will connect to the RPC endpoint, subscribe to new blocks, and print decoded transactions for the contract at `0x28172273CC1E0395F3473EC6eD062B6fdFb15940`.
