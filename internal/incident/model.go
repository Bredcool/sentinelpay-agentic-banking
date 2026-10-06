package incident

import "time"

type Status string

const (
	StatusDetected         Status = "detected"
	StatusInvestigating    Status = "investigating"
	StatusDecision         Status = "decision"
	StatusAwaitingApproval Status = "awaiting_approval"
	StatusExecuting        Status = "executing"
	StatusVerifying        Status = "verifying"
	StatusResolved         Status = "resolved"
	StatusEscalated        Status = "escalated"
)

type FailureType string

const (
	FailureUnknown           FailureType = "unknown"
	FailureInsufficientFunds FailureType = "insufficient_funds"
	FailureCurrencyMismatch  FailureType = "currency_mismatch"
	FailureInvalidParameter  FailureType = "invalid_parameter"
	FailureTransient         FailureType = "transient"
	FailureDownstream        FailureType = "downstream_failure"
)

type Incident struct {
	ID          string      `json:"id"`
	PaymentID   string      `json:"payment_id"`
	Status      Status      `json:"status"`
	FailureType FailureType `json:"failure_type"`
	Confidence  float64     `json:"confidence"`

	Evidence       []Evidence     `json:"evidence"`
	Recommendation Recommendation `json:"recommendation"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Evidence struct {
	Source string `json:"source"`
	Key    string `json:"key"`
	Value  string `json:"value"`
}

type Recommendation struct {
	Action           string `json:"action"`
	Reason           string `json:"reason"`
	RequiresApproval bool   `json:"requires_approval"`
}
