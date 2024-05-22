package avalanche

import (
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"math/rand"
)

// Transaction represents a transaction in the network.
// It consists of a nonce and some arbitrary data.
type Transaction struct {
	Nonce uint64 // Nonce is a unique identifier for the transaction.
	Data  int32  // Data is the content of the transaction.
}

// RandomTransaction creates a new random transaction.
// It generates a transaction with a random nonce and random data.
// Returns a pointer to the newly created Transaction.
func RandomTransaction() *Transaction {
	return &Transaction{
		Nonce: rand.Uint64(),        // Generates a random 64-bit unsigned integer for Nonce.
		Data:  int32(rand.Intn(10)), // Generates a random integer in the range [0, 9] for Data.
	}
}

// Serialize serializes the transaction into a byte slice.
// It converts the Nonce field of the transaction into a byte slice using little-endian encoding.
// Returns a byte slice representing the serialized transaction.
func (t *Transaction) Serialize() []byte {
	buf := make([]byte, 8)                      // Creates a byte slice with a length of 8 bytes.
	binary.LittleEndian.PutUint64(buf, t.Nonce) // Encodes the Nonce into the byte slice using little-endian order.
	return buf                                  // Returns the byte slice containing the serialized Nonce.
}

// Hash calculates the hash of the transaction.
// It first serializes the transaction, then computes the SHA-256 hash of the serialized data.
// Returns a string representing the hexadecimal encoding of the hash.
func (t *Transaction) Hash() string {
	hash := sha256.Sum256(t.Serialize()) // Computes the SHA-256 hash of the serialized transaction.
	return hex.EncodeToString(hash[:])   // Converts the hash to a hexadecimal string and returns it.
}

// ConflictWith checks if the transaction conflicts with another transaction.
// Two transactions are considered conflicting if they have the same Nonce.
// Returns true if the transactions conflict, false otherwise.
func (t *Transaction) ConflictWith(other *Transaction) bool {
	return t.Nonce == other.Nonce // Returns true if the Nonce of both transactions are equal.
}
