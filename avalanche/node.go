package avalanche

import (
	"sync"
)

// Node represents a node in the network
type Node struct {
	sync.Mutex
	ID      int64
	Mempool map[string]*TxState
}

// NewNode creates a new Node
func NewNode(id int64) *Node {
	return &Node{
		ID:      id,
		Mempool: make(map[string]*TxState),
	}
}

// HandleMessage handles a message received by the node
func (n *Node) HandleMessage(origin int64, msg Message) {
	// Implementation of message handling, similar to Rust version
	// Example: Handle incoming transactions and update mempool
	switch m := msg.(type) {
	case MessageTransaction:
		txHash := m.Tx.Hash()
		n.Lock()
		defer n.Unlock()
		if _, exists := n.Mempool[txHash]; !exists {
			n.Mempool[txHash] = &TxState{}
		}
		// Further processing...
	}
}
