# Architectural Manifesto: Layered Architecture (ADD-48)

## 1. The Engineering Paradigm
Layered Architecture establishes a hierarchy of responsibility. The fundamental rule is the **Dependency Rule**: Source code dependencies can only point inwards. The inner layers (Domain/Service) must know nothing about the outer layers (Presentation/Database).

## 2. Core Layers
- 2.1. Presentation Layer (Transport): Handles HTTP/gRPC, JSON/Protobuf. It is the "I/O" entry point.
- 2.2. Service Layer (Domain/Application): The "Brain." Contains pure business logic. It operates on abstractions (Interfaces).
- 2.3. Repository/Data Layer: The "Infrastructure." Implementation details of persistence (SQL, NoSQL, Cache).

## 3. The Dependency Inversion Principle (DIP)
This is the heart of layered architecture. The Service Layer does NOT depend on the Repository concrete class. It depends on a Repository Interface. This allows us to swap the infrastructure (Database) without ever touching the business logic.

## 4. The "Strict" vs. "Relaxed" Rule
- Strict Layering: A layer can only access the layer directly below it. (Maximizes encapsulation).
- Relaxed Layering: A layer can access any layer beneath it. (Better performance, but higher coupling).
*Principal Decision:* We enforce Strict Layering for all core business domains to minimize side effects.
