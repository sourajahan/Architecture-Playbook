# 17. Documentation Strategy: Knowledge as a First-Class Citizen

## 1. The Global Engineering Philosophy
In distributed, high-scale engineering teams, **Asynchronous Communication** is the only way to scale. If information lives only in the heads of senior engineers or in private chats (Slack/Teams), you have created a "Knowledge Silo." Knowledge Silos are the primary cause of architectural fragility and developer burnout. 

We treat **Documentation as a First-Class Engineering Artifact**, equal in importance to production code.

## 2. The "Docs-as-Code" Ecosystem
We reject Word documents, Confluence pages that rot, or disconnected PDFs. All documentation must live in the Git repository alongside the code.
- **Markdown:** The universal language of documentation.
- **Version Control:** Docs are versioned with the code. If you roll back a release, the documentation rolls back with it.
- **Automated Generation:** We generate API docs (OpenAPI/Swagger) directly from source code to ensure they are never out of sync.

## 3. Mandatory Engineering Artifacts
Every non-trivial architectural change must produce three types of artifacts:

### A. ADRs (Architecture Decision Records)
The "History of Why." 
- ADRs document the context, the decision, and—most importantly—the *rejected alternatives*. 
- This prevents the team from repeating the same mistakes or questioning a decision 6 months later.

### B. RFCs (Request for Comments)
The "Future of What."
- Before writing a single line of code for a major feature, we write an RFC. 
- This document describes the proposed design, trade-offs, and implementation plan. 
- It is shared globally across the organization to solicit feedback. It turns "Architectural Decisions" into "Community Decisions."

### C. Visual Architecture (The C4 Model)
- We avoid "bloated UML" that no one understands. 
- We standardize on the **C4 Model** (Context, Containers, Components, Code) to visualize our system at different levels of abstraction. 
- We use **Mermaid.js** or **PlantUML** to generate diagrams *as code*, ensuring they are diff-able in Git.

## 4. The "Evergreen" Policy (Maintenance)
Documentation that is incorrect is worse than no documentation—it misleads engineers.
- **The Boy Scout Rule applies to Docs:** If you find a bug in the documentation, you are required to fix it in the same Pull Request (PR) where you fix the code.
- **Stale Docs are Bugs:** If documentation is outdated, it is treated as a critical "Technical Debt" bug. If a module cannot be documented, it should be simplified until it can.

## 5. Global Onboarding (DX - Developer Experience)
A great project can be understood by a new engineer in under one hour.
- **README.md (The Front Door):** Must clearly define: "What is this?", "Why does it exist?", and "How do I run it in 5 minutes?".
- **Runbooks:** For every microservice, there must be a "Runbook" documenting: How to monitor it, how to scale it, and how to debug it when it catches fire at 3 AM.

---
*Status: Established*
