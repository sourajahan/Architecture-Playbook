# Architectural Manifesto: Modern Firewalling & Network Segmentation (ADD-70)

## 1. The Engineering Paradigm
The "Castle and Moat" security model is obsolete. We operate on "Zero Trust." Every service, pod, and machine is treated as if it were on the public internet. We do not trust internal traffic. We enforce security policies at the L3/L4 (Network) and L7 (Application) layers simultaneously.

## 2. Core Architectural Pillars
- 2.1. Micro-segmentation: We isolate services into small, logically separated zones. Service A cannot talk to Service B unless explicitly allowed by a Network Policy.
- 2.2. Egress Filtering (The Principal's Secret): Most breaches happen because an infected internal service calls home to a C2 (Command & Control) server. We block ALL outbound traffic by default and whitelist only verified external endpoints.
- 2.3. Layer 7 WAF (Web Application Firewall): L3/L4 firewalls are blind to HTTP payloads. WAF is mandatory to inspect traffic for SQLi, XSS, and bot patterns.
- 2.4. Infrastructure-as-Code (IaC): Manual firewall configuration is a high-severity risk. All network policies must be version-controlled (Terraform/Crossplane/Kubernetes NetworkPolicy).

## 3. The Principal's Warning
- Complexity is a Vulnerability: If your firewall rules are thousands of lines long and "no one dares to touch them," you are already compromised. Simplify or re-architect.
- Shadow IT: Developers spinning up rogue services that bypass the perimeter are your biggest security hole. Centralize network policy enforcement.
- False Sense of Security: A firewall doesn't fix a buggy application. It only limits the blast radius. Defense-in-depth is the goal.
