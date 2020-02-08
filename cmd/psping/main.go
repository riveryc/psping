package main

import (
	"github.com/riveryc/psping/pkg/ping"
	"os"
)

var (
	network = "tcp"
)

func main() {
	var targetInput TargetInput = TargetInput(os.Args[1])
	destination := targetInput.NewDest()
	ping.TCPPing(destination.Hostname, destination.Port, destination.IPv4Addr, network)
}
