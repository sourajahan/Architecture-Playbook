# Architectural Manifesto: Engineering Excellence (SOLID, TDD, DDD)

## 1. Domain-Driven Design (DDD): The Strategic Core
Complexity is not in the code; it is in the business logic.
- Bounded Contexts: We explicitly isolate domains. A "User" in the Billing service MUST be a different model than a "User" in the Auth service.
- Ubiquitous Language: Code must mirror the domain expert's language. If the business calls it a "Subscription," the class name must be `Subscription`.
- Aggregate Roots: We protect business invariants. Only the Aggregate Root can mutate its state, ensuring the system never enters an invalid state.

## 2. SOLID: Structural Integrity
SOLID principles are our architectural guardrails for maintainability.
- Dependency Inversion Principle (DIP): The highest priority. Our Business Logic (Domain) must NEVER depend on Infrastructure (SQL, Kafka, API). Everything is hidden behind Interfaces.
- Single Responsibility (SRP): If a class has more than one reason to change, it is a liability.
- Open/Closed Principle: We extend logic via composition, not modification.

## 3. Test-Driven Development (TDD): The Design Process
TDD is not about coverage; it is about architecture design through consumption.
- Contract-First: We write tests to define the API behavior before implementing it.
- Decoupling Metric: If code is hard to test (requires extensive mocking or setup), the architecture is flawed. We refactor until the design is clean.
- Safety: Tests serve as the primary documentation for how components interact in a distributed system.
