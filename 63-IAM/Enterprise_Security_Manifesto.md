# Architectural Manifesto: Zero Trust Identity & Access Management (ADD-63)

## 1. Executive Summary
In a distributed microservices environment, the "Network Perimeter" (Firewall) is dead. Security must be identity-centric, not network-centric. We operate under the "Zero Trust" assumption: Never trust, always verify. All services are potentially exposed; therefore, every service-to-service call must carry proof of identity (JWS) and transport-level security (mTLS).

## 2. The Identity Stack (Standards-Based)
- 2.1. OIDC (OpenID Connect): The standard for Identity (Authentication). Do not implement custom login flows. Use OIDC providers (Auth0/Keycloak/Okta).
- 2.2. OAuth2: The standard for Access (Authorization). Scopes are the granularity of access.
- 2.3. JWT (JSON Web Token): The payload for Identity Propagation. It must be short-lived, signed, and encrypted (JWE if PII exists).
- 2.4. mTLS (Mutual TLS): The hardware-level guarantee that the client and server are who they claim to be, using internal CA-signed certificates.

## 3. Core Architectural Pillars (The Enterprise Reality)
- 3.1. Identity Propagation: When User A calls Service 1, and Service 1 calls Service 2, the User Identity MUST be propagated (via JWT 'sub' claim). Do not lose the user context.
- 3.2. Token Lifecycle Management: Access tokens must be volatile (< 15 mins). Refresh tokens must be handled in a hardened storage. Token revocation (Blacklisting) is mandatory for compromised sessions.
- 3.3. Policy as Code (OPA - Open Policy Agent): Authorization logic (RBAC/ABAC) must NOT be hardcoded in services. Decouple policy evaluation to a sidecar (OPA) for centralized governance.
- 3.4. Key Rotation: Signing keys (JWKS) must rotate automatically. Services must poll the JWKS endpoint, not cache public keys indefinitely.
- 3.5. Audience & Scopes: Every token must strictly define the 'aud' (Audience) to prevent token reuse across services. Scope-based access prevents privilege escalation.
- 3.6. Auditability & Non-repudiation: Every identity request must be logged with a unique trace-id for forensics. 
- 3.7. Threat Modeling: Mitigation against Replay Attacks (Nonces/Exp), Token Theft (Sender-Constrained Tokens/DPoP), and Key Exposure (HSMs/Vaults).
