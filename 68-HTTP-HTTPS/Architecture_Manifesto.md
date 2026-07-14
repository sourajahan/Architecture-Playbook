# Architectural Manifesto: HTTP/HTTPS Strategy (ADD-68)

## 1. The Engineering Paradigm
We do not view HTTP as a text-based protocol; we view it as a stream of binary frames (HTTP/2 & 3). The shift from HTTP/1.1 (Request/Response) to HTTP/3 (QUIC/UDP) is not about "features"—it is about solving the "Head-of-Line Blocking" problem that kills performance in unstable networks.

## 2. Core Architectural Pillars
- 2.1. Connection Management: We never rely on default timeouts. We enforce `Keep-Alive` and connection pooling to avoid the "TCP Handshake Tax" on every request.
- 2.2. TLS Termination & Offloading: We terminate TLS at the Edge (Load Balancer/Gateway). This offloads the expensive CPU-bound crypto-operations, leaving the internal network to focus on application logic.
- 2.3. HSTS & Perfect Forward Secrecy: We enforce strict transport security to ensure that even if the server key is compromised later, past traffic remains decrypted.
- 2.4. Protocol Selection:
    - Internal (Service-to-Service): HTTP/2 (gRPC) for multiplexing.
    - External (Public API): HTTP/1.1 or HTTP/3 (QUIC) depending on network reliability.

## 3. The Principal's Warning
- The "Lazy TLS" Fallacy: Relying on self-signed certs or weak ciphers inside the DC ("it's just internal traffic") is the fastest way to get compromised. Internal traffic must use mTLS.
- HTTP/1.1 Limitations: Never use HTTP/1.1 for high-concurrency inter-service communication. It forces serial processing per connection.
- Caching Headers: If your API is read-heavy, leverage `Cache-Control` and `ETags` at the HTTP layer, not just the database layer.
