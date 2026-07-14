# Architectural Manifesto: Data Exposure & API Strategy (ADD-53)

## 1. The Engineering Paradigm
- **REST:** Resource-Oriented. It treats data as a set of static endpoints. It is the backbone of the web and excels at **Caching** and **Simplicity**.
- **GraphQL:** Graph-Oriented. It treats data as a graph of interconnected entities. It excels at **Data Aggregation** and **UI-Driven Requirements**.

## 2. Decision Matrix
| Feature | REST | GraphQL |
| :--- | :--- | :--- |
| **Data Fetching** | Fixed payloads (often Over/Under-fetching) | Client-defined (Precise) |
| **Caching** | Native HTTP Caching (Simple) | Complex (Needs Persisted Queries/CDN help) |
| **Schema** | Implicit (Swagger/OpenAPI required) | Strongly Typed (Built-in) |
| **Network** | Multiple round-trips (N+1 problem) | Single request (Aggregated) |

## 3. The Principal's Warning
- **The GraphQL N+1 Problem:** GraphQL is brilliant at avoiding the "Client-side N+1", but it creates a massive "Server-side N+1" if resolvers aren't architected with `DataLoader` (Batching/Caching).
- **The REST Overhead:** In mobile environments, REST often leads to multiple sequential requests, which kills battery and performance.
