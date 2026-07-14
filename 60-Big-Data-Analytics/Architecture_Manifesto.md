# Architectural Manifesto: Analytics Strategy (ADD-60)

## 1. The Engineering Paradigm
- **Hadoop/MapReduce (The Past):** Disk-bound, sequential, heavy I/O overhead. It was necessary when RAM was expensive. It is now a legacy architectural pattern.
- **Spark (The Standard):** Memory-bound, DAG-based (Directed Acyclic Graph) execution. It keeps the processing pipeline in RAM, resulting in ~100x performance gains.
- **The Principal's Shift:** We move from "Cluster-Centric" (managing nodes) to "Job-Centric" (Serverless/Managed Analytics).

## 2. Core Architectural Pillars
- 2.1. DAG (Directed Acyclic Graph): Spark optimizes the entire pipeline *before* executing a single line of code. It merges stages to reduce shuffling.
- 2.2. Decoupled Architecture: Storage (S3/GCS/Azure Blob) is independent of Compute (Spark Cluster/Serverless). If the cluster dies, the data remains.
- 2.3. Lazy Evaluation: Nothing happens until an "Action" (like `write` or `collect`) is triggered. This allows the framework to plan and prune data access.
