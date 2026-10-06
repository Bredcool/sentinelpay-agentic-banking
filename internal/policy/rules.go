package policy

type DefaultPolicy struct{}

func NewDefaultPolicy() *DefaultPolicy {
	return &DefaultPolicy{}
}

func (p *DefaultPolicy) Evaluate(
	action Action,
	failureType string,
	confidence float64,
) Decision {

	if confidence < 0.80 {
		return Decision{
			Allowed:          false,
			RequiresApproval: true,
			Action:           ActionEscalate,
			Reason:           "Agent confidence is below the automatic-action threshold",
		}
	}

	switch action {
	case ActionRetry:
		if failureType == "transient" {
			return Decision{
				Allowed: true,
				Action:  ActionRetry,
				Reason:  "Transient failure is eligible for bounded retry",
			}
		}

		return Decision{
			Allowed:          false,
			RequiresApproval: true,
			Action:           ActionEscalate,
			Reason:           "Failure type is not eligible for automatic retry",
		}

	default:
		return Decision{
			Allowed:          false,
			RequiresApproval: true,
			Action:           ActionEscalate,
			Reason:           "Action requires explicit policy approval",
		}
	}
}
