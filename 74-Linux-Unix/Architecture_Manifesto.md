# Architectural Manifesto: High-Performance Linux/Unix Systems Engineering (ADD-74)

## 1. Executive Summary
Every application we deploy is ultimately bounded by the limits of the POSIX subsystem and the Linux Kernel. As Principal Engineers, we do not treat the OS as a black box. We optimize process lifecycle, memory-mapped I/O, network socket buffers, and kernel parameterization to ensure maximum throughput and sub-millisecond latencies under peak enterprise workloads.

## 2. Core Architectural Pillars (The OS Internals)

### 2.1. The "Everything is a File" Philosophy (File Descriptors)
- In UNIX, sockets, pipes, devices, and files are all managed via File Descriptors (FDs).
- The default limit of 1024 FDs per process is a legacy bottleneck. Every high-concurrency connection occupies an FD. We programmatically and behaviorally bump limits using `RLIMIT_NOFILE` and configure `/etc/security/limits.conf` for system-wide resilience.

### 2.2. Memory Management & The OOM Killer
- **Virtual Memory & Overcommit:** Linux allows processes to allocate more memory than physically exists (`vm.overcommit_memory`). When memory runs out, the kernel invokes the Out-Of-Memory (OOM) Killer.
- **OOM Score:** The kernel calculates `oom_score` based on memory usage. Our mission-critical stateful systems (like databases/brokers) MUST have adjusted `oom_score_adj` to prevent the OS from killing them under memory pressure.
- **Page Cache & Direct I/O:** Reading from disk goes through the Kernel Page Cache. For high-performance databases, we bypass this cache using `O_DIRECT` or leverage memory-mapped files (`mmap`) to write directly to Virtual Memory.

### 2.3. The Evolution of I/O Multiplexing
- **Blocking I/O:** One thread per connection. Does not scale.
- **Non-Blocking + Epoll (L7 Gateways/Nginx/Go Runtime):** Event-driven notification system. Scale to millions of concurrent connections by allowing a single thread to monitor thousands of FDs.
- **io_uring (Modern Linux 5.1+):** True asynchronous I/O that completely eliminates the user-space/kernel-space context-switching overhead for I/O operations. This is the future of low-latency architecture.

### 2.4. Processes, Lightweight Threads (LWP), and Isolation
- The Linux kernel does not distinguish between processes and threads; both are task structures created via the `clone()` system call.
- **Namespaces & Cgroups:** The foundation of Docker and Kubernetes. Namespaces isolate resources (Network, Process ID, Mounts, Users), while Control Groups (Cgroups) enforce hard limits on CPU, Memory, and Disk I/O to prevent "noisy neighbor" syndrome.

---

## 3. Kernel Parameterization (Strategic Production sysctl.conf)

For high-throughput distributed systems, we enforce the following system-level configurations:

```ini
# Maximize System File Descriptor Limits
fs.file-max = 2097152

# Maximize Network Socket Listen Backlog (Prevents connection drops under burst)
net.core.somaxconn = 65535

# Set TCP Buffer Sizes (Empowering high-bandwidth, high-latency links)
net.ipv4.tcp_rmem = 4096 87380 16777216
net.ipv4.tcp_wmem = 4096 65536 16777216

# Enable TCP BBR Congestion Control (Google's highly optimized algorithm)
net.core.default_qdisc = fq
net.ipv4.tcp_congestion_control = bbr

# Prevent Swapping to Disk (Prioritize keeping application memory in physical RAM)
vm.swappiness = 10
