package avalanche

// Message is an interface representing a generic message in the network.
// It acts as a base interface for different types of messages that can be
// transmitted over the network. Implementations of this interface should
// provide the specific structure and data pertinent to the type of message
// being transmitted.
type Message interface{}

// MessageTransaction is a struct that implements the Message interface,
// representing a message specifically containing a transaction. This
// structure is used to encapsulate transaction data within a message format
// that can be transmitted over the network.
//
// Fields:
// Tx: A pointer to a Transaction object. The Transaction object contains all
//
//	the necessary data and operations associated with a transaction. By
//	using a pointer, the MessageTransaction can efficiently reference the
//	transaction data without copying the entire transaction structure.
//
// Example Usage:
// tx := &Transaction{...}  // Assume Transaction is a previously defined struct
// msg := MessageTransaction{Tx: tx}
// sendMessage(msg)  // sendMessage is a function that sends a Message over the network
type MessageTransaction struct {
	Tx *Transaction
}
