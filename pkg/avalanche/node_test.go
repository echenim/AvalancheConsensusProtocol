package avalanche

import (
	"testing"
)

func TestNewNode(t *testing.T) {
	tests := []struct {
		id   int64
		want int64
	}{
		{id: 1, want: 1},
		{id: 2, want: 2},
		{id: 3, want: 3},
	}

	for _, tt := range tests {
		t.Run("TestNewNode", func(t *testing.T) {
			node := NewNode(tt.id)
			if node.ID != tt.want {
				t.Errorf("NewNode(%d) = %d, want %d", tt.id, node.ID, tt.want)
			}
		})
	}
}

// TODO: bug to fix in this test
func TestHandleMessage(t *testing.T) {
	node := NewNode(1)

	tx1 := RandomTransaction()
	tx2 := RandomTransaction()
	tx3 := RandomTransaction()

	tests := []struct {
		origin   int64
		msg      Message
		expected int
	}{
		{origin: 2, msg: MessageTransaction{Tx: tx1}, expected: 1},
		{origin: 2, msg: MessageTransaction{Tx: tx1}, expected: 2},
		{origin: 3, msg: MessageTransaction{Tx: tx2}, expected: 1},
		{origin: 4, msg: MessageTransaction{Tx: tx3}, expected: 1},
		{origin: 4, msg: MessageTransaction{Tx: tx3}, expected: 2},
	}

	for _, tt := range tests {
		t.Run("TestHandleMessage", func(t *testing.T) {
			node.HandleMessage(tt.origin, tt.msg)
			txHash := tt.msg.(MessageTransaction).Tx.Hash()
			node.Lock()
			txState, exists := node.Mempool[txHash]
			node.Unlock()
			if !exists {
				t.Errorf("Transaction %s not found in Mempool", txHash)
			} else if txState.Confidence != tt.expected {
				t.Errorf("Transaction %s vote count = %d, want %d", txHash, txState.Confidence, tt.expected)
			}
		})
	}
}

func TestQueryNode(t *testing.T) {
	node1 := NewNode(1)
	node2 := NewNode(2)

	tx := RandomTransaction()
	txHash := tx.Hash()

	node1.Mempool[txHash] = NewTxState(tx)
	node1.Mempool[txHash].Preference = true

	tests := []struct {
		peer     *Node
		txHash   string
		expected bool
	}{
		{peer: node1, txHash: txHash, expected: true},
		{peer: node2, txHash: txHash, expected: false},
	}

	for _, tt := range tests {
		t.Run("TestQueryNode", func(t *testing.T) {
			result := node1.QueryNode(tt.peer, tt.txHash)
			if result != tt.expected {
				t.Errorf("QueryNode() = %v, want %v", result, tt.expected)
			}
		})
	}
}
