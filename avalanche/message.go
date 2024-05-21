package avalanche

// Message represents a message in the network
type Message interface{}

// MessageTransaction is a message containing a transaction
type MessageTransaction struct {
	Tx *Transaction
}
