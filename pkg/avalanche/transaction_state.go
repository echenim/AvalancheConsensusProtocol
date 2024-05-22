package avalanche

// TxState represents the state of a transaction in a node.
// It keeps track of various metrics and flags that help determine
// the acceptance or rejection of a transaction based on a confidence mechanism.
type TxState struct {
	Tx                  *Transaction // The transaction being tracked.
	Confidence          int          // The confidence level of the transaction.
	Accepted            bool         // Flag indicating if the transaction has been accepted.
	Rejected            bool         // Flag indicating if the transaction has been rejected.
	ConfidenceThreshold int          // The threshold at which the transaction is accepted or rejected.
	Alpha               int          // Parameter for Snowflake consensus.
	Beta                int          // Parameter for Snowball consensus.
	SnowflakeCounter    int          // Counter for Snowflake mechanism.
	SnowballCounter     int          // Counter for Snowball mechanism.
	Preference          bool         // Preference flag for consensus.
}

// NewTxState creates a new TxState instance for a given transaction.
// Initializes all metrics and flags to their default values.
func NewTxState(tx *Transaction) *TxState {
	return &TxState{
		Tx:                  tx,
		Confidence:          0,
		Accepted:            false,
		Rejected:            false,
		ConfidenceThreshold: 10, // Default threshold for acceptance/rejection.
		Alpha:               1,  // Default Snowflake parameter.
		Beta:                5,  // Default Snowball parameter.
		SnowflakeCounter:    0,
		SnowballCounter:     0,
		Preference:          true,
	}
}

// IncrementVote increases the confidence level of the transaction.
// If the confidence level reaches the threshold, the transaction is accepted.
// Uses Snowflake and Snowball consensus mechanisms to determine acceptance.
func (ts *TxState) IncrementVote() {
	if !ts.Accepted && !ts.Rejected {
		ts.Confidence++
		ts.SnowflakeCounter++
		if ts.Confidence >= ts.ConfidenceThreshold {
			ts.Accepted = true
		}
		if ts.SnowflakeCounter >= ts.Beta {
			ts.SnowballCounter++
			ts.SnowflakeCounter = 0
		}
		if ts.SnowballCounter >= ts.Beta {
			ts.Accepted = true
		}
	}
}

// DecrementVote decreases the confidence level of the transaction.
// If the confidence level falls below the negative threshold, the transaction is rejected.
// Resets the Snowflake counter on disagreement.
func (ts *TxState) DecrementVote() {
	if !ts.Accepted && !ts.Rejected {
		ts.Confidence--
		ts.SnowflakeCounter = 0 // Reset Snowflake counter on disagreement.
		if ts.Confidence <= -ts.ConfidenceThreshold {
			ts.Rejected = true
		}
	}
}
