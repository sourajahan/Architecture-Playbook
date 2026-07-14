# Architectural Manifesto: PKI & Automated Trust (ADD-64)

## 1. The Engineering Paradigm
PKI is the backbone of Zero Trust. It converts "Identity" into a cryptographic guarantee. We do not use self-signed certificates in production; we use a tiered CA hierarchy (Root -> Intermediate -> Leaf). 

## 2. Core Architectural Pillars
- 2.1. Root of Trust: The Root CA must be offline, air-gapped, and stored in an HSM (Hardware Security Module). It never signs leaf certificates.
- 2.2. Short-Lived Certificates (The Security Standard): Certificates should not live longer than 24-72 hours. This eliminates the need for massive CRLs (Certificate Revocation Lists).
- 2.3. Automated Lifecycle (ACME/Vault): We use the ACME protocol or Vault PKI Secrets Engine to automate issuance and renewal. If a human touches a cert, it's a failure.
- 2.4. Trust Stores: Every service must have a centralized Trust Store injected into its container/pod, ensuring it only trusts our Internal CA.

## 3. The Principal's Warning
- Outage Risk: An expired root or intermediate certificate is an "Atomic Bomb" for your infrastructure. All service-to-service communication will stop instantly.
- Monitoring: You MUST alert on cert expiration with 30-day, 14-day, and 7-day triggers.
- HSM/KMS: Never store the CA private keys in disk or environment variables. Use Vault, AWS KMS, or Google Cloud KMS.
