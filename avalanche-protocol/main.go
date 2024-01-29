package main

import (
	"fmt"
	"math/rand"
	"time"

	ap "github.com/echenim/consensus/avalanche-protocol/avalanche"
)

func main() {
	net := ap.NewNetwork(10)
	net.Run()

	for {
		tx := ap.RandomTransaction()
		fmt.Printf("sending new transaction into the network %s\n", tx.Hash())

		// Pick a random node in the network let the node handle the random transaction
		id := rand.Int63n(int64(len(net.nodes)))
		node := net.nodes[id]
		node.Lock()
		node.HandleMessage(0, &ap.MessageTransaction{tx})
		node.Unlock()

		time.Sleep(500 * time.Millisecond)
	}
}
