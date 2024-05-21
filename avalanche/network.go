package avalanche

import "time"

// Network represents the network of nodes
type Network struct {
	nodes map[int64]*Node
}

// NewNetwork creates a new network with n nodes
func NewNetwork(n int64) *Network {
	nodes := make(map[int64]*Node)
	for i := int64(0); i < n; i++ {
		nodes[i] = NewNode(i)
	}
	return &Network{nodes: nodes}
}

// Run starts the network operation
// Run starts the network operation
func (n *Network) Run() {
	for {
		for _, node := range n.nodes {
			n.simulateNodeActivity(node)
		}
		time.Sleep(1 * time.Second)
	}
}

func (n *Network) simulateNodeActivity(node *Node) {
	// Node broadcasts its transactions to random peers
	node.Lock()
	defer node.Unlock()

	for _, txState := range node.Mempool {
		// Simulate broadcasting the transaction to other nodes
		for _, peer := range n.getRandomPeers(node.ID, 3) {
			peer.HandleMessage(node.ID, MessageTransaction{Tx: txState.Tx})
		}
	}
}

func (n *Network) getRandomPeers(nodeID int64, count int) []*Node {
	peers := []*Node{}
	for _, node := range n.nodes {
		if node.ID != nodeID {
			peers = append(peers, node)
			if len(peers) >= count {
				break
			}
		}
	}
	return peers
}
