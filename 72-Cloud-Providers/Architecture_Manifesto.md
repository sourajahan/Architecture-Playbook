# Architectural Manifesto: Cloud Agnosticism & Value Extraction (ADD-72)

## 1. The Engineering Paradigm
We do not build a "Cloud-Agnostic" system by default, as it adds 30-50% engineering overhead. Instead, we build an **Abstraction Layer**. We use high-level managed services (PaaS/SaaS) for speed, but isolate them behind defined interfaces (Adapters/Ports). If we ever need to migrate, we only rewrite the adapter, not the business logic.

## 2. Core Architectural Pillars
- 2.1. FinOps as Architecture: The Cloud bill is a reflection of your architecture. If it costs too much, your architecture is wrong. We enforce "Cost per Transaction" metrics.
- 2.2. Data Gravity: Move compute to the data. If your Big Data is in GCS, your Spark jobs should run in GCP. Do not fight gravity.
- 2.3. Managed vs. Unmanaged: If it's not our core competency (e.g., managing a database or a load balancer), we buy a managed service (SaaS/PaaS). If it's our core IP, we control it (IaaS).
- 2.4. Infrastructure Abstraction: We use Terraform/Crossplane to manage resources, not the Console (ClickOps).

## 3. The Principal's Warning
- The "Multi-Cloud" Trap: Don't build for multi-cloud unless your scale requires it. It doubles the operational complexity, security surface, and team training requirements.
- API Overuse: Avoid "Serverless Vendor Lock-in" where you rely on proprietary triggers (e.g., DynamoDB Streams). Use standard interfaces (e.g., Kafka/PubSub) where possible.
