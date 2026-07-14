# Architectural Manifesto: OWASP Top 10 & Security-as-Code (ADD-64)

## 1. The Engineering Paradigm
We do not react to vulnerabilities; we eliminate "classes of vulnerabilities" by architectural design. OWASP Top 10 is our absolute baseline, not a target. Security is a property of the system, not a feature we add on top.

## 2. Core Architectural Pillars
- 2.1. Shift-Left (IDE Integration): Security linting (Semgrep/Snyk) must run on the developer's machine *before* code is committed.
- 2.2. Supply Chain Security (SCA): 80% of code is open source. You must automate Software Composition Analysis to block known CVEs in dependencies.
- 2.3. Threat Modeling: Before writing a single line of code, we ask: "How would an attacker abuse this endpoint?"
- 2.4. Defense in Depth: If the WAF fails, the API Gateway must block it; if the Gateway fails, the Application Code must validate the input.

## 3. The Principal's Warning
- Manual Testing is a Failure: If you rely on periodic penetration tests to find bugs, you have already lost.
- OWASP is not enough: OWASP Top 10 is just the "known" threats. Your architecture must also defend against business-logic-level exploits (e.g., IDOR, race conditions) that scanners will never find.
