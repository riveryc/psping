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
	//fmt.Println(destination.Hostname)
	//fmt.Println(destination.Port)
	//fmt.Println(destination.HasPort)
	ping.TCPPing(destination, pingLimit, network)
}
