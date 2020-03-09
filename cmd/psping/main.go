package main

import (
	"github.com/riveryc/psping/pkg/ping"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

var (
	network = "tcp"
	ipv4    = &cli.BoolFlag{
		Name:    "ipv4",
		Aliases: []string{"4"},
		Usage:   "Force to use IPv4",
	}
	ipv6 = &cli.BoolFlag{
		Name:    "ipv6",
		Aliases: []string{"6"},
		Usage:   "Force to use IPv6",
	}
)

func main() {
	//os.Args = []string{"psping", "-4", "google.com:80"}
	app := &cli.App{
		Name:    "psping",
		Usage:   "implements Ping functionality, TCP ping, latency and bandwidth measurement. ",
		Version: "v0.0.1",
		Flags: []cli.Flag{
			ipv4,
			ipv6,
		},
		Action: func(c *cli.Context) error {
			var targetInput TargetInput = TargetInput(c.Args().Get(0))
			destination := targetInput.NewDest()
			ping.TCPPing(destination.Hostname, destination.Port, destination.IPv4Addr, network)
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
