# 29. Git: The Architecture of Collaboration & History

## 1. The Engineering Paradigm
Git is the "Source of Truth" for our engineering culture. A clean, descriptive history is not a luxury; it is a necessity for debugging, auditing, and on-boarding new engineers. We treat commit messages as technical documentation.

## 2. Structural Principles
- **Branching Strategy:** We prefer **Trunk-Based Development** for high-velocity teams. Long-lived feature branches are an architectural anti-pattern that creates "Merge Hell."
- **Commit Hygiene (Conventional Commits):** Every commit must follow a semantic structure: `type(scope): description`.
    - `feat:` New functionality.
    - `fix:` Bug resolution.
    - `docs:` Documentation updates.
    - `refactor:` Code changes without functional impact.
- **Signed Commits:** In enterprise environments, code integrity is non-negotiable. All commits must be GPG/SSH signed to verify the author's identity.

## 3. The "Gatekeeper" Architecture
- **Pre-Commit Hooks:** Never allow "dirty" code to enter the repository. Automated checks (Linting, Formatting, Secret scanning) must run locally before a commit is even created.
- **Branch Protection:** No code enters the trunk without Peer Review and passing CI/CD pipelines. This is our architectural safety net.
- **Atomic Commits:** One commit = One logical change. Never mix unrelated bug fixes and feature additions in a single commit.

## 4. Operational Excellence
- **Rebase vs. Merge:** Use `rebase` locally to keep your history linear and clean. Use `merge --no-ff` (no-fast-forward) when merging feature branches to preserve the history of integration.
