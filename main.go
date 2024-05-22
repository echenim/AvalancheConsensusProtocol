package main

import (
	"fmt"
	"math/rand"

	"github.com/echenim/consensus/pkg/avalanche"
)

func main() {
	// Create a new network with 10 nodes
	network := avalanche.NewNetwork(10)

	// Generate a random transaction
	tx := avalanche.RandomTransaction()

	// Print the transaction details
	fmt.Printf("Transaction: %+v\n", tx)
	fmt.Printf("Transaction Hash: %s\n", tx.Hash())

	// Run the network in a separate goroutine
	go network.Run()

	// Simulate sending a transaction to a node
	node := network.Nodes[0]
	node.HandleMessage(1, avalanche.MessageTransaction{Tx: tx})

	// Generate conflicting transaction
	conflictingTx := &avalanche.Transaction{
		Nonce: tx.Nonce,
		Data:  int32(rand.Intn(10)),
	}
	node.HandleMessage(1, avalanche.MessageTransaction{Tx: conflictingTx})

	// Keep the program running
	select {}
}
