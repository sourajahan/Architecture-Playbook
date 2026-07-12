# 32. Trello: SDLC Workflow Engine

## Architectural Principle: Traceable Execution
Trello serves as the visualization layer for our SDLC. We enforce strict lifecycle stages: [Backlog -> In Progress -> Code Review -> QA -> Done].

## Design Rules:
1. **Automation-First:** Never create cards manually for recurring bugs or system incidents. Let the system create them.
2. **State Syncing:** Use Webhooks to sync Trello card status with Git branches. If a branch is merged, the card should move to "Code Review" automatically.
3. **Definition of Done (DoD):** Attach checklists to cards via API to enforce compliance requirements before a task can move to "Done."
