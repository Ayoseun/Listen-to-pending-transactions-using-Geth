package main

import (
	"context"
	"fmt"
	"log"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient/gethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

func main() {
	// Connect to an Ethereum node using an HTTP-based RPC client
	client, err := rpc.Dial("wss://polygon-mumbai.infura.io/ws/v3/b2eaf08a40e040a9b0a9947867492626")
	if err != nil {
		log.Fatal(err)
	}

	// Create a context to handle the subscription
	ctx := context.Background()

	// Subscribe to pending transactions
	 pendingTxSub := make(chan common.Hash)
	sub, err := gethclient.New(client).SubscribePendingTransactions(ctx,pendingTxSub)
	if err != nil {
		log.Fatal(err)
	}
	defer sub.Unsubscribe()

	// Start processing incoming pending transactions
	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case txs := <-pendingTxSub:
			fmt.Printf("Pending Transaction Hash: %s\n", txs.Hex())
			
		}
	}
}
