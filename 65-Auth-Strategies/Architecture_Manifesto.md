# Architectural Manifesto: Identity & Access Strategy (ADD-65)

## 1. The Engineering Paradigm
We do not build custom authentication protocols. We implement industry standards (OIDC/OAuth2). The focus shifts from "How do I check the password?" to "How do I propagate identity across 100 microservices without creating a bottleneck?"

## 2. Core Architectural Pillars
- 2.1. Centralized Identity, Decentralized Enforcement: The ID Provider (IdP) is central, but the Validation (AuthN/AuthZ) is decentralized (at the Edge or Sidecar).
- 2.2. Token-Based Identity (Stateless): We use JWTs for service-to-service communication. They must be signed, short-lived, and contain the minimum necessary claims.
- 2.3. Policy-Based Access Control (PBAC): We do not hardcode `if user.role == 'admin'` in the business logic. We delegate authorization to an engine like OPA (Open Policy Agent).
- 2.4. Identity Propagation: The user context MUST follow the request across the call graph (via `X-User-ID` or JWT propagation).

## 3. The Principal's Warning
- Don't Build: If you are building your own "Auth Server" from scratch, you are likely failing the project. Use Keycloak, Auth0, Okta, or AWS Cognito.
- Revocation Dilemma: Stateless JWTs are impossible to revoke instantly. If your business requires instant revocation, use "Opaque Tokens" (Reference tokens) with a distributed cache (Redis) check for every request (or keep JWTs extremely short, < 5min).
