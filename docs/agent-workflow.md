```markdown
# Agent Workflow

SentinelPay uses an observe → reason → act → verify loop.

## 1. Observe

Collect payment, account, and transaction context.

## 2. Reason

Classify the failure and identify the most likely root cause.

## 3. Decide

Select a recovery strategy.

## 4. Policy Check

Evaluate whether the proposed action is permitted.

## 5. Act

Execute the approved action through the Airwallex integration.

## 6. Verify

Check whether the expected state transition actually occurred.

## 7. Escalate

If verification fails or the action is outside the agent's authority, stop and escalate.