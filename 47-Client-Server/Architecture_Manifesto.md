# Architectural Manifesto: Contract-First Client/Server Communication (ADD-47)

## 1. The Engineering Paradigm
Client/Server is the atomic unit of all distributed systems. Our primary goal is to minimize the "Network Penalty." We prioritize reducing round-trips and minimizing payload size.

## 2. Core Architectural Pillars
- 2.1. Contract-First Design: Before writing code, we define the API schema (gRPC/Protobuf or OpenAPI/JSON). This allows the Client and Server teams to work in parallel.
- 2.2. Statelessness: The server must not assume any client state. All context must be passed via tokens (JWT) or headers. This enables horizontal scaling.
- 2.3. Resilient Communication: Networks will fail. Every client request must implement Retries (with exponential backoff), Timeouts, and Circuit Breaking.

## 3. Principal's Warning (Common Anti-Patterns)
- N+1 Query Problem: Making multiple sequential round-trips to the server for dependent data. We use batching or GraphQL to solve this.
- Chatty APIs: Sending too many small requests instead of one large, optimized payload.
- Versioning Hell: Modifying the API without backward compatibility. Always use API versioning (v1, v2) or header-based routing.
