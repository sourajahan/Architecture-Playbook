# Architectural Manifesto: Messaging & Event-Driven Topology (ADD-55)

## 1. The Engineering Paradigm
A Messaging Queue is not just a pipe; it is a "buffer" that allows us to absorb spikes in traffic without crashing our downstream services. We move from "Synchronous Requests" (Request/Response) to "Asynchronous Events" (Fire-and-Forget).

## 2. Core Architectural Pillars
- 2.1. Temporal Decoupling: Producer and Consumer do not need to be online at the same time. This is the ultimate resilience pattern.
- 2.2. Load Leveling (Spike Absorption): We size our services for *average* load, not *peak* load. The queue absorbs the burst.
- 2.3. Delivery Guarantees:
    - At-most-once: Fast, but risky (data loss).
    - At-least-once: Reliable, but requires Idempotency (data duplication).
    - Exactly-once: The holy grail (extremely expensive, usually at the broker level).

## 3. The Principal's Warning (Common Pitfalls)
- Poison Pills: A message that crashes the consumer. If not handled, it creates an infinite retry loop that blocks the entire queue.
- Ordering Guarantees: Most people think they need strict ordering. Most don't. Strict ordering severely limits throughput and partitioning scalability.
- Eventual Consistency: Switching to queues means giving up ACID transactions across services. You MUST handle data inconsistency in the business logic (Saga Pattern).
