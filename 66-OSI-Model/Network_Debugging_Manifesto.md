# Architectural Manifesto: OSI Debugging & Abstraction Strategy (ADD-66)

## 1. The Engineering Paradigm
The OSI model is a tool for "Divide and Conquer." Every architectural problem in distributed systems is a symptom of a layer failure.
- L1-L2 (Physical/Data Link): Rarely our problem (Cloud provider's job).
- L3 (Network): Routing loops, MTU mismatches, VPC Peering issues.
- L4 (Transport/TCP): The most common silent killer. TCP Handshakes, Congestion Windows, Retransmissions, Head-of-Line blocking.
- L7 (Application/HTTP/gRPC): Business logic, Serialization, Auth-layer overhead.

## 2. Abstraction Leaks
The most dangerous architecture is one that ignores "Leaky Abstractions." 
- If a gRPC service (L7) is slow, it might be because TCP (L4) has a small receive window.
- If a REST API is dropping connections, it might be an IP Table (L3) limit on the Load Balancer.
- We must maintain visibility at every layer. If the telemetry stops at L7, you are flying blind.

## 3. Core Architectural Pillars for Resiliency
- 3.1. Layered Observability: We need metrics for L3 (ICMP latency), L4 (TCP Retransmission count), and L7 (HTTP 5xx rates).
- 3.2. MTU Awareness: In containerized environments (Kubernetes/Calico), MTU mismatches between Pods and Nodes often lead to silent packet drops.
- 3.3. Connection Pooling vs. Keep-Alive: Poorly configured L4 Keep-Alive settings cause "Ghost Connections" that waste server resources.
- 3.4. Load Balancing Strategies: L4 Load Balancing (IP/Port) is fast but dumb. L7 Load Balancing (URL/Header-based) is slow but intelligent. Use L4 for high-throughput ingress, L7 for routing logic.
