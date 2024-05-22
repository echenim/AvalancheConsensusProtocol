package avalanche

import (
	"time"
)

// Network represents a network of nodes where each node can participate in the network's operations.
type Network struct {
	Nodes map[int64]*Node // A map of node IDs to their corresponding Node objects.
}

// NewNetwork creates a new network with a specified number of nodes.
// It initializes a map of nodes, where each node is assigned a unique ID from 0 to n-1.
func NewNetwork(n int64) *Network {
	nodes := make(map[int64]*Node)  // Initialize the map to store nodes.
	for i := int64(0); i < n; i++ { // Loop to create n nodes.
		nodes[i] = NewNode(i) // Create a new node with ID i and add it to the map.
	}
	return &Network{Nodes: nodes} // Return the newly created network.
}

// Run starts the network's operation loop.
// In this loop, each node in the network simulates its activity periodically.
func (n *Network) Run() {
	for {
		for _, node := range n.Nodes { // Iterate through all nodes in the network.
			n.simulateNodeActivity(node) // Simulate activity for each node.
		}
		time.Sleep(1 * time.Second) // Sleep for 1 second before the next iteration.
	}
}

// simulateNodeActivity simulates the activity of a given node in the network.
// It locks the node, processes transactions in the node's mempool, and updates their states based on peer queries.
func (n *Network) simulateNodeActivity(node *Node) {
	node.Lock()         // Lock the node to ensure thread-safe access.
	defer node.Unlock() // Unlock the node when the function completes.

	for txHash, txState := range node.Mempool { // Iterate through the transactions in the node's mempool.
		if txState.Accepted || txState.Rejected { // Skip transactions that have already been accepted or rejected.
			continue
		}
		preference := n.queryPeers(node, txHash) // Query peers to determine their preference for the transaction.
		if preference {
			txState.IncrementVote() // Increment the vote count if the peers prefer the transaction.
		} else {
			txState.DecrementVote() // Decrement the vote count if the peers do not prefer the transaction.
		}
	}
}

// queryPeers queries a random set of peers to determine their preference for a specific transaction.
// It returns true if the majority of the queried peers prefer the transaction.
func (n *Network) queryPeers(node *Node, txHash string) bool {
	votes := 0
	peers := n.getRandomPeers(node.ID, 5) // Get a random set of 5 peers.
	for _, peer := range peers {
		if peer.QueryNode(node, txHash) { // Query each peer for their preference.
			votes++
		}
	}
	return votes >= len(peers)/2 // Return true if the majority of peers prefer the transaction.
}

// getRandomPeers returns a slice of randomly selected peers, excluding the node with the specified nodeID.
// The number of peers returned is limited by the count parameter.
func (n *Network) getRandomPeers(nodeID int64, count int) []*Node {
	peers := []*Node{} // Initialize a slice to store the selected peers.
	for _, node := range n.Nodes {
		if node.ID != nodeID { // Exclude the node with the specified ID.
			peers = append(peers, node) // Add the node to the peers slice.
			if len(peers) >= count {    // Stop if the required number of peers is reached.
				break
			}
		}
	}
	return peers // Return the selected peers.
}
