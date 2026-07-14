# Architectural Manifesto: gRPC Strategy (ADD-51)

## 1. The Engineering Paradigm
gRPC is the gold standard for internal service communication. It operates on HTTP/2 (multiplexing) and uses Protocol Buffers (binary serialization). The primary value proposition is **Strong Typing** and **Schema Evolution**.

## 2. Core Architectural Pillars
- 2.1. Contract-First: We define the API via `.proto` files. This is the single source of truth. If it's not in the proto file, it doesn't exist.
- 2.2. Binary Serialization (Protobuf): JSON is slow and bulky (text-based). Protobuf is binary, faster to serialize/deserialize, and significantly smaller in payload.
- 2.3. HTTP/2 Multiplexing: Unlike HTTP/1.1 (which requires multiple TCP connections), HTTP/2 allows multiple requests to be sent over a *single* TCP connection, eliminating Head-of-Line blocking.

## 3. The Principal's Decision Matrix
- **Use gRPC for:** Internal Microservices, High-performance streaming, Polyglot environments (auto-generated client code in any language).
- **Avoid gRPC for:** Public-facing APIs (use REST/GraphQL instead), Browser-based clients (requires gRPC-Web proxy).
