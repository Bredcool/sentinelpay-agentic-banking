package policy

type Action string

const (
	ActionRetry    Action = "retry"
	ActionCancel   Action = "cancel"
	ActionEscalate Action = "escalate"
	ActionNoAction Action = "no_action"
)

type Decision struct {
	Allowed          bool
	RequiresApproval bool
	Action           Action
	Reason           string
}

type Policy interface {
	Evaluate(
		action Action,
		failureType string,
		confidence float64,
	) Decision
}
