package main

import (
	"github.com/riveryc/psping/pkg/ping"
	"math"
	"os"
)

var (
	network   = "tcp"
	pingLimit = math.MaxInt64
)

func main() {
	var targetInput TargetInput = TargetInput(os.Args[1])
	destination := targetInput.NewDest()
	ping.TCPPing(destination.Hostname, destination.Port, destination.IPv4Addr, network, pingLimit)
}
