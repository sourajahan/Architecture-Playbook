# Architectural Manifesto: From Legacy ETL to Modern ELT (ADD-59)

## 1. The Engineering Paradigm
- ETL (Extract, Transform, Load): The legacy approach. Transform data *before* loading. Bottleneck: Requires pre-processing servers and limits re-usability of raw data.
- ELT (Extract, Load, Transform): The modern approach. Load raw data into the Warehouse (or Lake) first, then transform using the Warehouse's massive parallel compute (e.g., BigQuery/Snowflake).

## 2. Core Architectural Pillars
- 2.1. Decoupled Compute & Storage: In modern warehouses, storage is cheap (Cloud Object Store), and compute scales independently.
- 2.2. Dimensional Modeling: Even in the cloud era, the **Star Schema** (Fact + Dimension tables) remains the fastest way to serve BI queries.
- 2.3. Idempotency: Data pipelines must be idempotent. If a job fails and we re-run it, we should not get duplicate records (using Upsert/Merge patterns).

## 3. The Principal's Decision Matrix
- **Data Warehouse:** Best for structured, curated, "source of truth" business reporting.
- **Data Lake:** Best for raw, unorganized, exploratory data (Schema-on-read).
- **Data Lakehouse:** The convergence (e.g., Delta Lake/Iceberg). The future.
