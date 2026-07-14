# Architectural Manifesto: The NoSQL Paradigm (ADD-57)

## 1. The Engineering Paradigm
We do not choose NoSQL because it's "fast." We choose it when the relational model becomes a bottleneck for a specific workload. We embrace **Polyglot Persistence**: The right tool for the right data access pattern.

## 2. The CAP Theorem Trade-off
In distributed systems, we follow the CAP Theorem:
$$ \text{Consistency} + \text{Availability} + \text{Partition Tolerance} = \text{Only 2} $$
NoSQL systems generally prioritize **Availability** and **Partition Tolerance** (AP) over strict **Consistency** (CP).

## 3. Categorization by Access Pattern
- **Key-Value (e.g., Redis, DynamoDB):** Optimized for ultra-fast lookups ($O(1)$).
- **Document (e.g., MongoDB):** Optimized for nested, polymorphic data.
- **Wide-Column (e.g., Cassandra):** Optimized for massive write-throughput and time-series.
- **Graph (e.g., Neo4j):** Optimized for deep relationship traversal (social networks/fraud detection).

## 4. The Principal's Warning
- **Denormalization:** In NoSQL, you model data based on the **Query**, not the object. If you need to join 5 tables, you aren't using NoSQL properly; you should have embedded that data.
- **Eventual Consistency:** You must write your application code to handle the fact that "Data might not be there yet" after a write.
