package agent

import "context"

type InvestigationRequest struct {
	IncidentID string
	PaymentID  string
}

type InvestigationResult struct {
	FailureType string
	Confidence  float64
	Reason      string
	Evidence    []Evidence
}

type Evidence struct {
	Source string
	Key    string
	Value  string
}

type Agent interface {
	Investigate(
		ctx context.Context,
		req InvestigationRequest,
	) (InvestigationResult, error)
}
