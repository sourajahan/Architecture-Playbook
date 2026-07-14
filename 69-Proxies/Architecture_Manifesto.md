# Architectural Manifesto: Proxy & Gateway Strategy (ADD-69)

## 1. The Engineering Paradigm
The Proxy is no longer just a router. It is the "Data Plane" of your architecture. We distinguish between:
- **Forward Proxy:** Controls outbound traffic (Egress). Used for security/compliance (Filtering/Auditing).
- **Reverse Proxy:** Controls inbound traffic (Ingress). Used for Load Balancing, SSL Termination, and Caching.
- **Sidecar Proxy (Service Mesh):** The modern evolution. Every microservice has a local proxy (e.g., Envoy) managing mTLS, Retries, and Telemetry transparently.

## 2. Core Architectural Pillars
- 2.1. Separation of Concerns: Application code should NEVER handle TLS, Retries, Circuit Breaking, or Request Retries. The Proxy handles this.
- 2.2. Programmability: We choose proxies that are programmable (e.g., Envoy, Nginx with Lua, or custom Go/Rust proxies).
- 2.3. Observability at the Edge: The proxy is the single source of truth for request latency, error rates, and traffic patterns.

## 3. The Principal's Warning
- Avoid "Proxy Hell": Do not chain proxies unnecessarily (e.g., Client -> Load Balancer -> Ingress Gateway -> API Gateway -> Sidecar -> App). Every hop adds latency and debugging complexity.
- Offloading: If your application is doing TLS termination, Gzip compression, or Auth validation, you are wasting CPU cycles. Move these to the Proxy/Gateway layer.
- Health Checks: The proxy must be the "Health Checker." If the proxy dies, the traffic stops. High availability (HA) for proxies is more critical than for the services themselves.
