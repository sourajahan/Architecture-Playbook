# Architectural Manifesto: Service Oriented Architecture (SOA)

## 1. The Engineering Paradigm
SOA is about creating an enterprise ecosystem where business logic is encapsulated in reusable services. Unlike Microservices, which prioritize team autonomy, SOA prioritizes **Enterprise Service Governance** and **Reusable Business Assets**.

## 2. Core Pillars of SOA
- 2.1. Service Contract: All communication is via explicit, versioned contracts (WSDL/Protobuf/OpenAPI). The implementation is hidden.
- 2.2. Service Governance: We maintain a Registry/Catalog of services. Any team can discover and consume an existing service instead of reinventing the wheel.
- 2.3. Orchestration vs. Choreography: In SOA, we often use an Orchestrator (Service Bus/Workflow Engine) to coordinate complex transactions across multiple shared services.

## 3. The Principal's Verdict (SOA vs. Microservices)
- Microservices: Optimization for Team Autonomy (Speed of change).
- SOA: Optimization for Business Reuse (Consistency and shared assets).
*Warning:* Don't build an "Enterprise Service Bus" (ESB) unless you have a massive legacy integration requirement. ESBs become a "God Object" and the primary bottleneck of your system.
