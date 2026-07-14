# CAP/PACELC Trade-off Manifesto  

## 1. The Physics of Distributed Systems
CAP dictates: In the presence of a Network Partition (P), you MUST trade off between Consistency (C) and Availability (A).

## 2. The PACELC Expansion
Since Partition (P) is a reality, PACELC adds: Else (E), when the system is running normally (no partition), do we trade Latency (L) for Consistency (C)?
- PC/EC: Consistency over Latency/Availability (e.g., Google Spanner).
- PA/EL: Availability over Consistency (e.g., Cassandra, DynamoDB).

## 3. Decision Matrix for Principal Engineers
- Use CP systems for Financial/Inventory records (where accuracy is paramount).
- Use AP systems for User profiles, Session stores, and Social feeds (where uptime is paramount).
