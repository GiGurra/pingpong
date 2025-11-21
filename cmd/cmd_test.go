package cmd

import (
	"context"
	"fmt"
	"net"
	"testing"
	"time"
)

func TestTCPPingPong(t *testing.T) {
	port, err := getFreePort()
	if err != nil {
		t.Fatalf("Failed to get free port: %v", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Start TCP server
	go func() {
		if err := runListen(ctx, port, ConnTypeTCP); err != nil {
			// If the test finishes and cancels context, runListen might return nil or error depending on timing.
			// We generally ignore errors here unless startup failed immediately.
		}
	}()

	// Wait for server to start
	if !waitForPort(port, "tcp", 2*time.Second) {
		t.Fatalf("Server failed to start on port %d", port)
	}

	// Run client
	if err := runPing(ctx, ConnTypeTCP, "localhost", port); err != nil {
		t.Fatalf("Ping failed: %v", err)
	}
}

func TestUDPPingPong(t *testing.T) {
	port, err := getFreePort()
	if err != nil {
		t.Fatalf("Failed to get free port: %v", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Start UDP server
	go func() {
		if err := runListen(ctx, port, ConnTypeUDP); err != nil {
			// Ignore
		}
	}()

	// Wait for server (UDP is connectionless, so we can't "connect" to check if it's listening easily without sending data)
	// But runListenUDP starts pretty fast.
	time.Sleep(200 * time.Millisecond)

	// Run client
	if err := runPing(ctx, ConnTypeUDP, "localhost", port); err != nil {
		t.Fatalf("Ping failed: %v", err)
	}
}

func getFreePort() (int, error) {
	l, err := net.Listen("tcp", ":0")
	if err != nil {
		return 0, err
	}
	defer l.Close()
	return l.Addr().(*net.TCPAddr).Port, nil
}

func waitForPort(port int, network string, timeout time.Duration) bool {
	deadline := time.Now().Add(timeout)
	for time.Now().Before(deadline) {
		conn, err := net.DialTimeout(network, net.JoinHostPort("localhost", fmt.Sprintf("%d", port)), 100*time.Millisecond)
		if err == nil {
			conn.Close()
			return true
		}
		time.Sleep(50 * time.Millisecond)
	}
	return false
}
