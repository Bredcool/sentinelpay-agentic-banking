# SentinelPay Architecture

SentinelPay follows a bounded agentic architecture.

```text
Payment Incident
       |
       v
Incident Manager
       |
       v
Agent Investigation
       |
       +----> Airwallex
       |
       v
Root Cause Analysis
       |
       v
Recovery Decision
       |
       v
Policy Engine
       |
   +---+---+
   |       |
 Safe    Approval
 Action   Required
   |       |
   +---+---+
       |
       v
Airwallex Action
       |
       v
Verification
       |
   +---+---+
   |       |
Resolved Escalated