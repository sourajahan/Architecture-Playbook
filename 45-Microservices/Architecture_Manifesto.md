# Architectural Manifesto: Microservices & Domain Isolation (ADD-45)

## 1. The Engineering Paradigm
Microservices is the architectural manifestation of Conway's Law. We decompose the system not by "function," but by "Business Capability" (Bounded Contexts). If you cannot deploy a service without deploying three others, you are failing the definition of Microservices.

## 2. Core Architectural Pillars
- 2.1. Service Autonomy: Each service must own its data and schema. No cross-service database access.
- 2.2. Distributed Resilience: We assume the network is unreliable. Every remote call must implement timeouts, retries, and circuit breakers.
- 2.3. Decentralized Governance: We allow polyglot persistence and technology stacks, provided the service adheres to the organization's API standards (e.g., gRPC/OpenAPI).

## 3. The Principal's Warning (Anti-Patterns)
- Distributed Monolith: Tight coupling via synchronous REST calls and shared databases. This is the worst of both worlds.
- Microservice Overkill: Breaking a system into services too early, leading to massive operational overhead (service discovery, tracing, deployment pipelines).
