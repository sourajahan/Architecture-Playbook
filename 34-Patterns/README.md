# 34. Advanced Presentation Patterns: State Orchestration & Coupling Metrics

## 1. The Architectural Challenge: State Mutation
The primary source of bugs in UI development is "Uncontrolled State Mutation." Whether using MVC, MVP, or MVVM, the goal is always the same: **Constraint of Side Effects.** 

- **MVC:** Often leads to "Fat Controllers." It creates a circular dependency between View and Controller, which is the #1 killer of unit-testability in UI layers.
- **MVP:** Introduced "Passive Views." It solves the tight coupling problem but leads to an explosion of interfaces (one for every View).
- **MVVM:** The optimal pattern for high-interactivity apps. By introducing a "ViewModel" as a transformation layer, we decouple the UI from the underlying business logic, enabling independent evolution of the Presentation layer.

## 2. Engineering Principles for Modern UI
- **Unidirectional Data Flow (UDF):** Regardless of the pattern, implement UDF. Data moves down (Store -> ViewModel -> View), Actions move up (View -> Action -> Store). This prevents "State Explosion."
- **Immutability as Default:** Never mutate objects in the UI layer. Treat state as a stream of snapshots.
- **Dependency Inversion (DI):** The View should never know about the concrete Implementation of the ViewModel. It should only know the Interface.

## 3. When to deviate from Patterns?
In modern "Component-Based" architectures (React/Vue/Svelte), these classic patterns often collapse into **"Functional Composition."** Do not force a pattern where a simple composition of pure functions serves the architecture better.
