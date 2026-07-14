# Architectural Manifesto: SQL Databases at Scale (ADD-56)

## 1. The Engineering Paradigm
SQL is the anchor of distributed systems. It provides the "Ground Truth." While NoSQL provides scale, SQL provides trust. Our strategy is to maximize the life of SQL using proper sharding, indexing, and replication strategies before considering NoSQL alternatives.

## 2. Core Architectural Pillars
- 2.1. Read/Write Splitting: We never run heavy analytics on the Primary (Master) node. We use Read Replicas for scalability.
- 2.2. Connection Pooling: SQL connections are expensive. We manage them via proxies (e.g., PgBouncer) or middleware, never creating them on-the-fly.
- 2.3. Horizontal Scaling (Sharding): When vertical scaling (bigger CPU/RAM) reaches its limit, we shard the data (partitioning by user_id or tenant_id).
- 2.4. Indexing Strategy: An index is a "trade-off." It speeds up Reads (SELECT) but slows down Writes (INSERT/UPDATE). We design indexes based on query execution plans, not intuition.

## 3. The Principal's Decision Matrix
- When to stay on SQL: If the schema is structured, consistency is non-negotiable, and complex joins are frequent.
- When to move to NoSQL: If the write-volume exceeds the capability of a single write-master, or the data model is highly polymorphic (schema-less).
