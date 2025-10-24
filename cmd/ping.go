package cmd

import (
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
			switch params.ConnType {
			case ConnTypeTCP:
				conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", params.Addr, params.Port))
				if err != nil {
					log.Fatalf("Failed to connect to %s:%d - %v", params.Addr, params.Port, err)
				}
				defer func() {
					err := conn.Close()
					if err != nil {
						log.Fatalf("Failed to close connection - %v", err)
					}
				}()
				fmt.Printf("Successfully connected to %s:%d via TCP\n", params.Addr, params.Port)
				_, err = conn.Write([]byte("ping"))
				if err != nil {
					log.Fatalf("Failed to send ping - %v", err)
				}
				buf := make([]byte, 1024)
				n, err := conn.Read(buf)
				if err != nil {
					log.Fatalf("Failed to read response - %v", err)
				}
				slog.Info("Received response", "response", string(buf[:n]))
			case ConnTypeUDP:
				conn, err := net.Dial("udp", fmt.Sprintf("%s:%d", params.Addr, params.Port))
				if err != nil {
					log.Fatalf("Failed to connect to %s:%d - %v", params.Addr, params.Port, err)
				}
				defer func() {
					err := conn.Close()
					if err != nil {
						log.Fatalf("Failed to close connection - %v", err)
					}
				}()
				fmt.Printf("Successfully connected to %s:%d via UDP\n", params.Addr, params.Port)
				_, err = conn.Write([]byte("ping"))
				if err != nil {
					log.Fatalf("Failed to send ping - %v", err)
				}
				buf := make([]byte, 1024)
				n, err := conn.Read(buf)
				if err != nil {
					log.Fatalf("Failed to read response - %v", err)
				}
				slog.Info("Received response", "response", string(buf[:n]))
			default:
				log.Fatalf("Unknown connection type: %s", params.ConnType)
			}
		},
	}
}
