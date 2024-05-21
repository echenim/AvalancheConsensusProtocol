package avalanche

import (
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"math/rand"
)

// Transaction represents a transaction in the network
type Transaction struct {
	Nonce uint64
	Data  int32
}

// RandomTransaction creates a new random transaction
func RandomTransaction() *Transaction {
	return &Transaction{
		Nonce: rand.Uint64(),
		Data:  int32(rand.Intn(10)),
	}
}

// Serialize serializes the transaction
func (t *Transaction) Serialize() []byte {
	buf := make([]byte, 8)
	binary.LittleEndian.PutUint64(buf, t.Nonce)
	return buf
}

// Hash calculates the hash of the transaction
func (t *Transaction) Hash() string {
	hash := sha256.Sum256(t.Serialize())
	return hex.EncodeToString(hash[:])
}