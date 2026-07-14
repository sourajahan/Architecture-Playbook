package main

import (
	"fmt"
	"net"
	"syscall"
)

// ConfigureHighPerformanceSocket applies socket-level tuning for high-throughput services.
func ConfigureHighPerformanceSocket(conn net.Conn) error {
	// Cast to raw file descriptor to apply OS-level tweaks
	tcpConn, ok := conn.(*net.TCPConn)
	if !ok {
		return fmt.Errorf("not a TCP connection")
	}

	rawConn, err := tcpConn.SyscallConn()
	if err != nil {
		return err
	}

	return rawConn.Control(func(fd uintptr) {
		// 1. Set Read/Write Buffer Size (Scale for high-bandwidth)
		// Default is often 128KB, we might need 1MB+ for high-latency/high-bandwidth pipes.
		bufferSize := 1024 * 1024 
		syscall.SetsockoptInt(int(fd), syscall.SOL_SOCKET, syscall.SO_RCVBUF, bufferSize)
		syscall.SetsockoptInt(int(fd), syscall.SOL_SOCKET, syscall.SO_SNDBUF, bufferSize)

		// 2. TCP_NODELAY (Disable Nagle's algorithm)
		// Essential for RPC/gRPC where we need minimum latency, not packet bundling.
		syscall.SetsockoptInt(int(fd), syscall.IPPROTO_TCP, syscall.TCP_NODELAY, 1)
		
		// 3. TCP_QUICKACK (Optional: Send ACKs immediately)
		// Caution: Can increase load on the sender.
		syscall.SetsockoptInt(int(fd), syscall.IPPROTO_TCP, syscall.TCP_QUICKACK, 1)
	})
}
