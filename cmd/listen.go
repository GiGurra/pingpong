package cmd

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net"

	"github.com/GiGurra/boa/pkg/boa"
	"github.com/spf13/cobra"
)

type ListenParams struct {
	ConnType ConnType `alts:"tcp,udp"`
	Port     int
}

func ListenCmd() boa.CmdIfc {
	return &boa.CmdT[ListenParams]{
		Use:         "listen",
		Long:        "listen/act as a server",
		ParamEnrich: boa.ParamEnricherDefault,
		RunFunc: func(params *ListenParams, cmd *cobra.Command, args []string) {
			if err := runListen(cmd.Context(), params.Port, params.ConnType); err != nil {
				log.Fatalf("Listen failed: %v", err)
			}
		},
	}
}

func runListen(ctx context.Context, port int, connType ConnType) error {
	switch connType {
	case ConnTypeTCP:
		return runListenTCP(ctx, port)
	case ConnTypeUDP:
		return runListenUDP(ctx, port)
	default:
		return fmt.Errorf("unknown connection type: %s", connType)
	}
}

func runListenTCP(ctx context.Context, port int) error {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return fmt.Errorf("failed to listen on port %d: %w", port, err)
	}

	go func() {
		<-ctx.Done()
		listener.Close()
	}()
	defer listener.Close()

	fmt.Printf("Listening on port %d via TCP\n", port)
	for {
		conn, err := listener.Accept()
		if err != nil {
			if ctx.Err() != nil {
				return nil
			}
			slog.Error("Failed to accept connection", "error", err)
			continue // Should we exit? Maybe. But original code continued (though it didn't have context).
			// Actually if listener is closed, Accept returns error. We should probably check if it's a closed network connection error.
			// But checking ctx.Err() is safer.
		}
		go handleTCPConnection(conn)
	}
}

func runListenUDP(ctx context.Context, port int) error {
	addr := net.UDPAddr{
		Port: port,
		IP:   net.ParseIP("0.0.0.0"),
	}
	conn, err := net.ListenUDP("udp", &addr)
	if err != nil {
		return fmt.Errorf("failed to listen on port %d: %w", port, err)
	}

	go func() {
		<-ctx.Done()
		conn.Close()
	}()
	defer conn.Close()

	fmt.Printf("Listening on port %d via UDP\n", port)
	buf := make([]byte, 1024)
	for {
		n, remoteAddr, err := conn.ReadFromUDP(buf)
		if err != nil {
			if ctx.Err() != nil {
				return nil
			}
			slog.Error("Failed to read from UDP", "error", err)
			continue
		}
		message := string(buf[:n])
		slog.Info("Received UDP message", "from", remoteAddr, "message", message)
		if message == "ping" {
			_, err = conn.WriteToUDP([]byte("pong"), remoteAddr)
			if err != nil {
				slog.Error("Failed to send pong", "error", err)
			} else {
				slog.Info("Sent pong response", "to", remoteAddr)
			}
		}
	}
}

func handleTCPConnection(conn net.Conn) {
	defer func() {
		err := conn.Close()
		if err != nil {
			slog.Error("Failed to close connection", "error", err)
		}
	}()
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		slog.Error("Failed to read from connection", "error", err)
		return
	}
	message := string(buf[:n])
	slog.Info("Received TCP message", "from", conn.RemoteAddr(), "message", message)
	if message == "ping" {
		_, err = conn.Write([]byte("pong"))
		if err != nil {
			slog.Error("Failed to send pong", "error", err)
		} else {
			slog.Info("Sent pong response", "to", conn.RemoteAddr())
		}
	}
}