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

type PingParams struct {
	ConnType ConnType `alts:"tcp,udp"`
	Addr     string
	Port     int
}

func PingCmd() boa.CmdIfc {
	return &boa.CmdT[PingParams]{
		Use:         "ping",
		Long:        "Ping a server",
		ParamEnrich: boa.ParamEnricherDefault,
		RunFunc: func(params *PingParams, cmd *cobra.Command, args []string) {
			if err := runPing(cmd.Context(), params.ConnType, params.Addr, params.Port); err != nil {
				log.Fatalf("Ping failed: %v", err)
			}
		},
	}
}

func runPing(ctx context.Context, connType ConnType, addr string, port int) error {
	var d net.Dialer
	address := fmt.Sprintf("%s:%d", addr, port)

	switch connType {
	case ConnTypeTCP:
		conn, err := d.DialContext(ctx, "tcp", address)
		if err != nil {
			return fmt.Errorf("failed to connect to %s via TCP: %w", address, err)
		}
		defer func() {
			if err := conn.Close(); err != nil {
				log.Printf("Failed to close connection: %v", err)
			}
		}()

		fmt.Printf("Successfully connected to %s via TCP\n", address)
		_, err = conn.Write([]byte("ping"))
		if err != nil {
			return fmt.Errorf("failed to send ping: %w", err)
		}

		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			return fmt.Errorf("failed to read response: %w", err)
		}
		slog.Info("Received response", "response", string(buf[:n]))
		return nil

	case ConnTypeUDP:
		conn, err := d.DialContext(ctx, "udp", address)
		if err != nil {
			return fmt.Errorf("failed to connect to %s via UDP: %w", address, err)
		}
		defer func() {
			if err := conn.Close(); err != nil {
				log.Printf("Failed to close connection: %v", err)
			}
		}()

		fmt.Printf("Successfully connected to %s via UDP\n", address)
		_, err = conn.Write([]byte("ping"))
		if err != nil {
			return fmt.Errorf("failed to send ping: %w", err)
		}

		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			return fmt.Errorf("failed to read response: %w", err)
		}
		slog.Info("Received response", "response", string(buf[:n]))
		return nil

	default:
		return fmt.Errorf("unknown connection type: %s", connType)
	}
}