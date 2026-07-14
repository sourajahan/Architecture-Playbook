# Architectural Manifesto: Enterprise Service Bus (ESB) & SOAP (ADD-52)

## 1. The Engineering Paradigm
SOAP and ESB represent the "Era of Heavyweight Integration." SOAP provides a formal, strictly-typed, XML-based contract (WSDL). The ESB acts as the "Centralized Orchestrator" that handles complex message transformation, routing, and protocol mediation.

## 2. Core Pillars of Legacy Integration
- 2.1. WSDL (Web Services Description Language): The formal contract. Unlike REST, it defines the operation, data types, and protocol in one rigid file.
- 2.2. Message Mediation: ESB is the "God Object" that can translate between protocols (e.g., JMS to SOAP) and formats (XML to JSON).
- 2.3. Transactional Integrity: SOAP/WS-AtomicTransaction was designed for distributed transactions (ACID over the wire), something modern REST APIs struggle with.

## 3. The Principal's Warning (The "ESB Trap")
- The God Object Anti-Pattern: Most ESBs eventually become the "Single Point of Failure." They become too complex to update and too risky to delete.
- Modernizing: We do not perform "Big Bang" migrations. We apply the **Strangler Fig Pattern**—slowly replacing ESB functions with Microservices while keeping the integration alive.
