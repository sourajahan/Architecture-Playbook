# Architectural Manifesto: OOP at Scale & The Actor Model (ADD-36)

## 1. Abstract
This document outlines our adoption of OOP as an encapsulation barrier and the Actor Model as our primary concurrency paradigm. We explicitly reject Shared-Mutable-State in favor of isolation and asynchronous message passing.

## 2. Object-Oriented Principles (Scale Perspective)
OOP is not about inheritance hierarchies; it is about managing system complexity through:
- 2.1. Encapsulation of Invariants: Objects must protect their internal state. No external actor can mutate an object’s state without going through defined domain methods.
- 2.2. Dependency Inversion: All infrastructure interactions (DB, Network) must be abstracted via Interfaces. The Business Logic (Domain) is agnostic of the underlying implementation.
- 2.3. Composition over Inheritance: We build functionality by composing interfaces, ensuring that behavior is injected at runtime, not bound at compile-time.

## 3. The Actor Model: Shared-Nothing Concurrency
Traditional threading relies on Mutexes/Locks, which are prone to deadlocks and performance degradation (Lock Contention). We adopt the Actor Model for the following reasons:
- 3.1. Isolation: Each Actor manages its own private state. No other entity can access or corrupt it.
- 3.2. Asynchronous Communication: Actors interact via 'Mailboxes' (Message Queues). This decouples the caller from the receiver.
- 3.3. Failure Domains: If an Actor crashes, it does not bring down the entire runtime. We implement 'Supervision Strategies' to restart failed actors without data loss.

## 4. Operational Requirements (The "Principal" Standard)
- 4.1. Non-blocking I/O: Actors must never perform blocking I/O calls. Every interaction must be offloaded to a dedicated context or worker pool.
- 4.2. Idempotency: Given the nature of message passing, message delivery guarantees (At-least-once) require that all Actor state transitions be idempotent.
- 4.3. Backpressure: Actors must monitor their Mailbox size. If an Actor is overwhelmed, it must signal the upstream producers to throttle traffic.

## 5. Anti-Patterns to Avoid
- 5.1. Shared Global State: Any object accessible by multiple actors is a violation of our architectural standard.
- 5.2. Direct Reference Leaks: Passing object pointers between actors is strictly forbidden. We pass immutable Data Transfer Objects (DTOs) only.
- 5.3. Deep Inheritance Chains: We limit inheritance depth to 1. Use composition (mixins/interfaces) for everything else.

## 6. Deployment & Observability
- 6.1. Tracing: Every message passed between actors must carry a 'TraceID' context for distributed observability (OpenTelemetry).
- 6.2. Supervision: We utilize parent-child actor supervision trees. If a child fails, the parent decides whether to resume, restart, or terminate the subtree.

## 7. Compliance
- Adherence to this manifesto is mandatory for all core infrastructure services. Exceptions must be documented in a formal ADR.
