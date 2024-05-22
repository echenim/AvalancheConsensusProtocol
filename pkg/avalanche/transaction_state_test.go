package avalanche

import (
	"testing"
)

//TODO: Bug to fix in the test
func TestTxState_IncrementVote(t *testing.T) {
	tests := []struct {
		name                     string
		initialConfidence        int
		initialAccepted          bool
		initialRejected          bool
		confidenceThreshold      int
		alpha                    int
		beta                     int
		expectedConfidence       int
		expectedAccepted         bool
		expectedRejected         bool
		expectedSnowflakeCounter int
		expectedSnowballCounter  int
	}{
		{
			name:                     "Increment below threshold",
			initialConfidence:        1,
			initialAccepted:          false,
			initialRejected:          false,
			confidenceThreshold:      10,
			alpha:                    1,
			beta:                     5,
			expectedConfidence:       2,
			expectedAccepted:         false,
			expectedRejected:         false,
			expectedSnowflakeCounter: 1,
			expectedSnowballCounter:  0,
		},
		// {
		// 	name:               "Increment to acceptance",
		// 	initialConfidence:  9,
		// 	initialAccepted:    false,
		// 	initialRejected:    false,
		// 	confidenceThreshold: 10,
		// 	alpha:              1,
		// 	beta:               5,
		// 	expectedConfidence: 10,
		// 	expectedAccepted:   true,
		// 	expectedRejected:   false,
		// 	expectedSnowflakeCounter: 10,
		// 	expectedSnowballCounter: 0,
		// },
		{
			name:                     "Increment to Snowball",
			initialConfidence:        9,
			initialAccepted:          false,
			initialRejected:          false,
			confidenceThreshold:      10,
			alpha:                    1,
			beta:                     5,
			expectedConfidence:       10,
			expectedAccepted:         true,
			expectedRejected:         false,
			expectedSnowflakeCounter: 1,
			expectedSnowballCounter:  1,
		},
		{
			name:                     "Increment to Snowball finalization",
			initialConfidence:        10,
			initialAccepted:          false,
			initialRejected:          false,
			confidenceThreshold:      10,
			alpha:                    1,
			beta:                     5,
			expectedConfidence:       11,
			expectedAccepted:         true,
			expectedRejected:         false,
			expectedSnowflakeCounter: 1,
			expectedSnowballCounter:  1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tx := &Transaction{Nonce: 1, Data: 1}
			ts := &TxState{
				Tx:                  tx,
				Confidence:          tt.initialConfidence,
				Accepted:            tt.initialAccepted,
				Rejected:            tt.initialRejected,
				ConfidenceThreshold: tt.confidenceThreshold,
				Alpha:               tt.alpha,
				Beta:                tt.beta,
				SnowflakeCounter:    tt.initialConfidence,
				SnowballCounter:     0,
				Preference:          true,
			}
			ts.IncrementVote()
			if ts.Confidence != tt.expectedConfidence {
				t.Errorf("expected confidence %d, got %d", tt.expectedConfidence, ts.Confidence)
			}
			if ts.Accepted != tt.expectedAccepted {
				t.Errorf("expected accepted %v, got %v", tt.expectedAccepted, ts.Accepted)
			}
			if ts.Rejected != tt.expectedRejected {
				t.Errorf("expected rejected %v, got %v", tt.expectedRejected, ts.Rejected)
			}
			if ts.SnowflakeCounter != tt.expectedSnowflakeCounter {
				t.Errorf("expected SnowflakeCounter %d, got %d", tt.expectedSnowflakeCounter, ts.SnowflakeCounter)
			}
			if ts.SnowballCounter != tt.expectedSnowballCounter {
				t.Errorf("expected SnowballCounter %d, got %d", tt.expectedSnowballCounter, ts.SnowballCounter)
			}
		})
	}
}

func TestTxState_DecrementVote(t *testing.T) {
	tests := []struct {
		name                string
		initialConfidence   int
		initialAccepted     bool
		initialRejected     bool
		confidenceThreshold int
		expectedConfidence  int
		expectedAccepted    bool
		expectedRejected    bool
	}{
		{
			name:                "Decrement above threshold",
			initialConfidence:   1,
			initialAccepted:     false,
			initialRejected:     false,
			confidenceThreshold: 10,
			expectedConfidence:  0,
			expectedAccepted:    false,
			expectedRejected:    false,
		},
		{
			name:                "Decrement to rejection",
			initialConfidence:   -9,
			initialAccepted:     false,
			initialRejected:     false,
			confidenceThreshold: 10,
			expectedConfidence:  -10,
			expectedAccepted:    false,
			expectedRejected:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tx := &Transaction{Nonce: 1, Data: 1}
			ts := &TxState{
				Tx:                  tx,
				Confidence:          tt.initialConfidence,
				Accepted:            tt.initialAccepted,
				Rejected:            tt.initialRejected,
				ConfidenceThreshold: tt.confidenceThreshold,
				Alpha:               1,
				Beta:                5,
				SnowflakeCounter:    0,
				SnowballCounter:     0,
				Preference:          true,
			}
			ts.DecrementVote()
			if ts.Confidence != tt.expectedConfidence {
				t.Errorf("expected confidence %d, got %d", tt.expectedConfidence, ts.Confidence)
			}
			if ts.Accepted != tt.expectedAccepted {
				t.Errorf("expected accepted %v, got %v", tt.expectedAccepted, ts.Accepted)
			}
			if ts.Rejected != tt.expectedRejected {
				t.Errorf("expected rejected %v, got %v", tt.expectedRejected, ts.Rejected)
			}
		})
	}
}
