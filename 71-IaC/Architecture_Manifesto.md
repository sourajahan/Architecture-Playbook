# Architectural Manifesto: Infrastructure as Code (ADD-71)

## 1. The Engineering Paradigm
We treat Infrastructure as Application Code. 
- **Immutable Infrastructure:** We never patch a running server. We destroy and recreate it from a fresh image.
- **Declarative vs. Imperative:** We define "What" the state should be, not "How" to change it.
- **State as Truth:** The "State File" (e.g., terraform.tfstate) is the most critical artifact. It must be versioned, encrypted, and locked.

## 2. Core Architectural Pillars
- 2.1. Idempotency: Running the same IaC code 100 times must result in the same outcome. No side effects.
- 2.2. Modularity: We build reusable, hardened modules (e.g., a "Golden VPC Module").
- 2.3. Automated Testing (TDD for Infra): We use tools like Terratest to validate infra before it touches production.
- 2.4. Drift Detection: We continuously scan for manual changes and overwrite them to match the code.

## 3. The Principal's Warning
- Manual Changes are Crimes: If an engineer makes a change in the Cloud Console (AWS/Azure) without updating the code, they have introduced "Configuration Drift." This leads to mysterious failures later.
- Secret Management: NEVER put secrets in your IaC code. Use Vault or Managed KMS.
