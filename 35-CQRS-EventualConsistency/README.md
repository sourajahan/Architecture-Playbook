# 35. CQRS & Eventual Consistency: Distributed Data Architecture

## 1. The Engineering Paradigm
CQRS segregates Read and Write workloads to optimize scaling. The Write side (Command) enforces domain invariants and consistency (ACID). The Read side (Query) serves materialized views, optimized for performance and specific UI/API requirements.

## 2. Core Architectural Pillars
*   **Command Side (Write Model):** Must be strictly consistent. This is the "Source of Truth."
*   **Query Side (Read Model):** A projection of the state. It is asynchronously updated, leading to Eventual Consistency.
*   **The Synchronization Gap:** We acknowledge that the Read side lags behind the Write side. We design systems that handle this latency (e.g., Optimistic UI, Caching).

## 3. The Dual-Write Problem & The Outbox Pattern
The most critical architectural risk is updating the Database and Publishing an Event (Kafka/RabbitMQ) separately. If one succeeds and the other fails, the system enters an inconsistent state.
*   **Solution:** The Transactional Outbox Pattern. We commit the State Change AND the Event to the same Database in a single Atomic Transaction. A separate Relay service then ensures delivery.
