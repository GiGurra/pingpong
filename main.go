package main

import (
	"github.com/GiGurra/boa/pkg/boa"
	"github.com/gigurra/pingpong/cmd"
)

func main() {
	boa.Cmd{
		Use:  "pingpong",
		Long: "Port forwarding tester",
		SubCmds: boa.SubCmds(
			cmd.PingCmd(),
			cmd.ListenCmd(),
		),
	}.Run()
}
