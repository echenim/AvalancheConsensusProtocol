package avalanche

import (
	"testing"
)

// TestNewNetwork tests the creation of a new network.
func TestNewNetwork(t *testing.T) {
	tests := []struct {
		name string
		n    int64
		want int64
	}{
		{"Zero nodes", 0, 0},
		{"One node", 1, 1},
		{"Ten nodes", 10, 10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			network := NewNetwork(tt.n)
			if int64(len(network.Nodes)) != tt.want {
				t.Errorf("NewNetwork(%d) = %d nodes, want %d nodes", tt.n, len(network.Nodes), tt.want)
			}
		})
	}
}

// TestGetRandomPeers tests the retrieval of random peers.
func TestGetRandomPeers(t *testing.T) {
	network := NewNetwork(10)

	tests := []struct {
		name   string
		nodeID int64
		count  int
		want   int
	}{
		{"Get 3 peers", 0, 3, 3},
		{"Get 5 peers", 1, 5, 5},
		{"Get more peers than available", 2, 20, 9}, // 9 because one node is excluded
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			peers := network.getRandomPeers(tt.nodeID, tt.count)
			if len(peers) != tt.want {
				t.Errorf("getRandomPeers(%d, %d) = %d peers, want %d peers", tt.nodeID, tt.count, len(peers), tt.want)
			}
		})
	}
}

// TestQueryPeers tests querying peers for their preference on a transaction.
func TestQueryPeers(t *testing.T) {
	network := NewNetwork(10)
	node := network.Nodes[0]
	tx := RandomTransaction()
	txHash := tx.Hash()
	node.Mempool[txHash] = NewTxState(tx)

	tests := []struct {
		name   string
		node   *Node
		txHash string
		want   bool
	}{
		{"Query peers with existing transaction", node, txHash, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := network.queryPeers(tt.node, tt.txHash)
			if got != tt.want {
				t.Errorf("queryPeers(%v, %s) = %v, want %v", tt.node, tt.txHash, got, tt.want)
			}
		})
	}
}

// TODO: test faling bug to fix
// TestSimulateNodeActivity tests the simulation of node activity.
func TestSimulateNodeActivity(t *testing.T) {
	network := NewNetwork(10)
	node := network.Nodes[0]
	tx := RandomTransaction()
	txHash := tx.Hash()
	node.Mempool[txHash] = NewTxState(tx)

	tests := []struct {
		name   string
		node   *Node
		txHash string
	}{
		{"Simulate activity with one transaction", node, txHash},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			network.simulateNodeActivity(tt.node)
			txState := tt.node.Mempool[tt.txHash]
			if txState.Confidence == 0 {
				t.Errorf("simulateNodeActivity did not update the transaction state")
			}
		})
	}
}
