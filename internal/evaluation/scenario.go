package evaluation

type Scenario struct {
	ID          string                 `json:"id"`
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	Input       map[string]interface{} `json:"input"`
	Expected    ExpectedOutcome        `json:"expected"`
}

type ExpectedOutcome struct {
	Status           string `json:"status"`
	Action           string `json:"action"`
	RequiresApproval bool   `json:"requires_approval"`
}
