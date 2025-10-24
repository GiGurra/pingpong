package cmd

import (
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
			switch params.ConnType {
			case ConnTypeTCP:
				listener, err := net.Listen("tcp", fmt.Sprintf(":%d", params.Port))
				if err != nil {
					log.Fatalf("Failed to listen on port %d - %v", params.Port, err)
				}
				defer func() {
					err := listener.Close()
					if err != nil {
						log.Fatalf("Failed to close listener - %v", err)
					}
				}()
				fmt.Printf("Listening on port %d via TCP\n", params.Port)
				for {
					conn, err := listener.Accept()
					if err != nil {
						slog.Error("Failed to accept connection", "error", err)
						continue
					}
					go handleTCPConnection(conn)
				}
			case ConnTypeUDP:
				addr := net.UDPAddr{
					Port: params.Port,
					IP:   net.ParseIP("0.0.0.0"),
				}
				conn, err := net.ListenUDP("udp", &addr)
				if err != nil {
					log.Fatalf("Failed to listen on port %d - %v", params.Port, err)
				}
				defer func() {
					err := conn.Close()
					if err != nil {
						log.Fatalf("Failed to close connection - %v", err)
					}
				}()
				fmt.Printf("Listening on port %d via UDP\n", params.Port)
				buf := make([]byte, 1024)
				for {
					n, remoteAddr, err := conn.ReadFromUDP(buf)
					if err != nil {
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
			default:
				log.Fatalf("Unknown connection type: %s", params.ConnType)
			}
		},
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
