package evaluation

type Result struct {
	ScenarioID string
	Passed     bool
	Failures   []string
}

type Evaluator struct{}

func (e *Evaluator) Evaluate(
	scenario Scenario,
	actual ExpectedOutcome,
) Result {

	var failures []string

	if actual.Status != scenario.Expected.Status {
		failures = append(
			failures,
			"unexpected final incident status",
		)
	}

	if actual.Action != scenario.Expected.Action {
		failures = append(
			failures,
			"unexpected recovery action",
		)
	}

	if actual.RequiresApproval != scenario.Expected.RequiresApproval {
		failures = append(
			failures,
			"unexpected approval requirement",
		)
	}

	return Result{
		ScenarioID: scenario.ID,
		Passed:     len(failures) == 0,
		Failures:   failures,
	}
}
