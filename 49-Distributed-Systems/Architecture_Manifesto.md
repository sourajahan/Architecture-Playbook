# Architectural Manifesto: Distributed Systems & The Fallacies of Computing (ADD-49)

## 1. The Engineering Paradigm
A distributed system is a collection of independent computers that appear to the user as a single coherent system. The fundamental challenge is not "adding servers," but **"maintaining coherence"** in the face of Partial Failure and Network Partition.

## 2. The 8 Fallacies of Distributed Computing (The Architect's Bible)
You must design as if the following are false:
1. The network is reliable.
2. Latency is zero.
3. Bandwidth is infinite.
4. The network is secure.
5. Topology doesn't change.
6. There is one administrator.
7. Transport cost is zero.
8. The network is homogeneous.

## 3. Core Architectural Pillars
- 3.1. Consensus Algorithms (Paxos/Raft): How do we agree on state when nodes are unreliable?
- 3.2. Observability (The Truth): If you can't trace a request across nodes, your system is a black box. You are blind.
- 3.3. Idempotency: Every distributed operation MUST be idempotent. If a retry fails, re-running the command must not corrupt the state.
- 3.4. Fail-Fast & Bulkheads: One failing service must not kill the entire cluster. Isolate failures using bulkheads (isolated thread pools/resources).
