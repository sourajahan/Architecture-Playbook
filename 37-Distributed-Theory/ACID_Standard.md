# ACID Architectural Standard (ADD-37)

## 1. The Paradigm
ACID defines the operational constraints for single-node transaction integrity. It is non-negotiable for financial and state-sensitive systems.

## 2. Core Pillars
- Atomicity: The "All-or-Nothing" property. If any part of the transaction fails, the entire transaction is rolled back.
- Consistency: Data must transition from one valid state to another, upholding all schema/business constraints.
- Isolation: Transactions must be executed as if they were serial. We use Isolation Levels (Read Committed, Serializable) to prevent dirty reads or lost updates.
- Durability: Once a transaction is committed, it survives system crashes. Write-Ahead Logging (WAL) is the standard mechanism here.

## 3. Engineering Constraint
Never assume ACID properties across distributed boundaries. ACID applies to the local Node/Database. For Distributed Systems, we look towards Sagas and Eventual Consistency (BASE).
