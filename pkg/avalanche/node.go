package avalanche

import (
	"sync"
)

// Node represents a node in the network.
// Each Node has a unique ID and a Mempool, which is a map to store the states of transactions.
// The sync.Mutex is embedded to provide thread-safe operations on the Node's data.
type Node struct {
	sync.Mutex
	ID      int64
	Mempool map[string]*TxState
}

// NewNode creates a new Node with a given ID.
// It initializes the Mempool as an empty map to store transaction states.
// Parameters:
// - id: an int64 representing the unique identifier for the node.
// Returns:
// - A pointer to a new Node.
func NewNode(id int64) *Node {
	return &Node{
		ID:      id,
		Mempool: make(map[string]*TxState),
	}
}

// HandleMessage processes a message received by the node.
// If the message is of type MessageTransaction, it extracts the transaction's hash and updates the Mempool.
// It ensures thread-safe operations by locking the Node's mutex before modifying the Mempool.
// Parameters:
// - origin: an int64 representing the ID of the node that sent the message.
// - msg: a Message interface which is type-asserted to MessageTransaction.
func (n *Node) HandleMessage(origin int64, msg Message) {
	switch m := msg.(type) {
	case MessageTransaction:
		txHash := m.Tx.Hash()
		n.Lock()
		defer n.Unlock()
		if txState, exists := n.Mempool[txHash]; !exists {
			// If the transaction is not in the Mempool, add a new TxState for it.
			n.Mempool[txHash] = NewTxState(m.Tx)
		} else {
			// If the transaction already exists in the Mempool, increment its vote count.
			txState.IncrementVote()
		}
	}
}

// QueryNode queries another node for their preference on a transaction with a given hash.
// It locks the peer node's mutex to safely access its Mempool.
// If the transaction state exists in the peer's Mempool, it returns the preference of the transaction.
// Parameters:
// - peer: a pointer to the Node being queried.
// - txHash: a string representing the hash of the transaction being queried.
// Returns:
// - A boolean indicating the peer node's preference for the transaction.
func (n *Node) QueryNode(peer *Node, txHash string) bool {
	peer.Lock()
	defer peer.Unlock()

	if txState, exists := peer.Mempool[txHash]; exists {
		return txState.Preference
	}
	return false
}
