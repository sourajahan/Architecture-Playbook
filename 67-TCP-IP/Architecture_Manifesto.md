# Architectural Manifesto: TCP/IP & Infrastructure Reality (ADD-67)

## 1. The Engineering Paradigm
While OSI provides the diagnostic vocabulary, TCP/IP provides the operational reality. In modern cloud environments (AWS/Azure/GCP), the Network Layer (IP) is virtualized (VPC/SDN), but the Transport Layer (TCP) remains the primary bottleneck for application performance.

## 2. OSI vs. TCP/IP (The Practical Mapping)
| OSI Layer | TCP/IP Layer | Operational Focus (Principal) |
| :--- | :--- | :--- |
| 7, 6, 5 | Application | Serialization, Auth, Load Balancing |
| 4 | Transport | Window Scaling, Congestion Control (BBR/CUBIC) |
| 3 | Internet | Routing, IP Fragmentation (MTU Issues) |
| 2, 1 | Link | MAC/Physical (Usually Cloud Provider/SDN) |

## 3. The Principal's Warning
- MTU (Maximum Transmission Unit): In containerized (Kubernetes) environments, packets larger than 1450 bytes (due to VXLAN/Overlay overhead) get silently dropped. If your service "hangs" on large requests but works on small ones, check your MTU.
- Congestion Control (BBR): If your services serve high-bandwidth streaming or large payloads, switch from the default CUBIC to Google's BBR congestion algorithm at the kernel level.
- TCP Handshake Overhead: For short-lived connections (REST), 3-way handshakes destroy performance. Use Persistent Connections (Keep-Alive) or HTTP/2/3 to reuse the socket.
