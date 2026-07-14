# Architectural Manifesto: Orchestration vs. Automation (ADD-54)

## 1. The Engineering Paradigm
- **BPM (The Discipline):** A methodology to model, analyze, and optimize business processes (usually via BPMN diagrams). It is the "What" and "Why."
- **BPEL (The Legacy Implementation):** An XML-based execution language from the SOA era. It attempted to standardize process logic. It failed because it was too rigid, XML-heavy, and hard to debug.
- **Modern Orchestration (The Reality):** We use Workflow-as-Code (like Temporal or Camunda) to manage stateful, long-running processes that span multiple microservices.

## 2. Core Pillars of Orchestration
- 2.1. Long-Running State: Microservices are stateless. Business processes are *stateful* (e.g., an order process taking 3 days). The Orchestrator manages this state.
- 2.2. Compensating Transactions (Saga Pattern): Since we don't have distributed ACID transactions, the orchestrator MUST know how to "undo" steps if a later step fails.
- 2.3. Observability: A business process is a "Black Box" if you cannot see the status of each step. The Orchestrator must expose the state of every process instance.

## 3. The Principal's Warning
- Avoid "God Orchestrators": Don't put too much business logic inside your Workflow engine. Keep logic in services; use the Orchestrator only for *coordination*.
- Avoid XML-BPEL: Unless you are maintaining 20-year-old banking legacy systems, do not start new projects with BPEL. It is technical debt by design.
