# 08. Requirements Elicitation

## Concept
Requirements Elicitation is the practice of researching, uncovering, and documenting the needs of stakeholders for a system. In software architecture, this is the most critical phase; an architect's job is to bridge the gap between business ambiguity and technical clarity.

## Core Principles
1. **Stakeholder Identification:** Distinguish between users, customers, and those who impact technical decisions.
2. **NFRs vs FRs:** Always prioritize Non-Functional Requirements (Performance, Scalability, Security) as they define the architecture.
3. **Ambiguity Resolution:** Never assume. If a requirement says "Support many users," define the scale (1k vs 1M concurrent users).

## Techniques Used
- **Stakeholder Interviews:** Direct feedback loop.
- **5 Whys Analysis:** Digging to the root cause.
- **Prototyping:** Validating assumptions early.
- **INVEST Criteria:** Ensuring user stories are Independent, Negotiable, Valuable, Estimable, Small, and Testable.

## Architectural Insight
"A requirement is not a request; it is a constraint." Every requirement you elicit directly influences your choice of patterns (e.g., Microservices vs. Monolith, SQL vs. NoSQL).

---
*Status: Completed*
