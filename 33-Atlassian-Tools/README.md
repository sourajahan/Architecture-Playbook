# 33. Atlassian Suite: Engineering Governance & SSOT

## 1. The Engineering Paradigm
Atlassian is the Operating System of Enterprise Engineering. Jira represents the **State Machine** of our delivery, and Confluence represents the **Single Source of Truth (SSOT)** of our architectural knowledge. We do not use them to "track work"; we use them to "enforce discipline."

## 2. Structural Principles
- **SSOT (Single Source of Truth):** If it is not in Confluence, it does not exist. We treat technical documentation as code. Outdated documentation is technical debt.
- **Workflow as Code:** Jira transitions should be strict. An issue should not move to "In Progress" without a linked branch, and not to "Done" without a successful CI/CD pipeline.
- **Data-Driven Delivery:** We use JQL (Jira Query Language) to measure velocity, cycle time, and bottlenecks. If we cannot measure the delivery process, we cannot improve it.

## 3. The "Principal" Deep-Dive
- **Jira Automation:** Use native automation rules to enforce policies. E.g., "If a ticket is idle for 5 days, auto-comment and ask for status." 
- **Documentation Hygiene:** Enforce templates in Confluence. Every ADR (Architectural Decision Record) must follow a standard structure: Context, Decision, Consequences.
- **Integration Layer:** Connect Jira, Bitbucket, and CI/CD tools. The engineering manager should be able to see the status of a feature from the Jira card without asking developers for updates.
