# Architectural Manifesto: Data Warehouse Design Principles (ADD-61)

## 1. The Engineering Paradigm
A Data Warehouse (DWH) is a "Single Source of Truth." If two departments report different numbers for "Revenue," the DWH has failed. We optimize for **Read-heavy workloads** using dimensional modeling.

## 2. Core Architectural Pillars
- 2.1. Defining the "Grain": Every fact table must have an atomic grain. If you cannot define what a single row represents (e.g., "one line item of one sale"), your model is flawed.
- 2.2. Dimensional Modeling (Star Schema): We prefer Star Schema over Snowflake. Denormalization reduces joins and improves performance for BI tools (Tableau/Looker).
- 2.3. Slowly Changing Dimensions (SCD): Historical data preservation is not optional. We use SCD Type 2 (History tracking) to know exactly what a customer's state was at the time of a transaction.

## 3. The Principal's Warning
- Avoid "Normalized Analytical Models": Forcing 3NF (Third Normal Form) on an analytical warehouse creates performance bottlenecks that require massive compute power to solve.
- Avoid "Black Box Transformations": If your ETL logic is hidden in proprietary GUI tools, you have no visibility. Everything must be code (dbt/SQL/Terraform).
