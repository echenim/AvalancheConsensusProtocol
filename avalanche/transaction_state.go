package avalanche

// TxState represents the state of a transaction in a node
type TxState struct {
	Tx                  *Transaction
	Confidence          int
	Accepted            bool
	Rejected            bool
	ConfidenceThreshold int
}

// NewTxState creates a new TxState
func NewTxState(tx *Transaction) *TxState {
	return &TxState{
		Tx:                  tx,
		Confidence:          0,
		Accepted:            false,
		Rejected:            false,
		ConfidenceThreshold: 10, // example threshold value
	}
}

// IncrementVote increases the confidence of the transaction
func (ts *TxState) IncrementVote() {
	if !ts.Accepted && !ts.Rejected {
		ts.Confidence++
		if ts.Confidence >= ts.ConfidenceThreshold {
			ts.Accepted = true
		}
	}
}

// DecrementVote decreases the confidence of the transaction
func (ts *TxState) DecrementVote() {
	if !ts.Accepted && !ts.Rejected {
		ts.Confidence--
		if ts.Confidence <= -ts.ConfidenceThreshold {
			ts.Rejected = true
		}
	}
}
