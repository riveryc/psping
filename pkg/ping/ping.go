package ping

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"runtime/pprof"
	"time"
)

func exitReport(hostname, port, network string, total, received *int) {
	// capture ctrl+c and stop CPU profiler
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for sig := range c {
			fmt.Printf("--- %s:%s %s ping statistics ---\n", hostname, port, network)
			failureRate := (*total - *received) * 100 / *total
			fmt.Printf("%d packets transmitted, %d packets received, %d %% packet loss\n", *total, *received, failureRate)
			log.Printf("captured %v, stopping profiler and exiting..", sig)
			pprof.StopCPUProfile()
			os.Exit(1)
		}
	}()
}

func TCPPing(hostname, port, IPv4Addr, network string) {
	fmt.Printf("PING %s:%s (%s):\n", hostname, port, IPv4Addr) //TODO: pick first IPV4 address
	r := 0
	i := 0
	exitReport(hostname, port, network, &i, &r)
	for {
		time.Sleep(1 * time.Second)
		i = i + 1
		now := time.Now()
		connDial, err := net.DialTimeout(network, hostname+":"+port, 5*time.Second)
		if err != nil {
			// handle error
			fmt.Println(err.Error())
			continue
		}
		localAddrPort := connDial.LocalAddr().String()
		remoteAddrPort := connDial.RemoteAddr().String()
		done := time.Now()
		latency := done.Sub(now)
		fmt.Printf("dial %s %s: from %s tcp_seq=%d time=%d ms\n", network, remoteAddrPort, localAddrPort, i, latency.Milliseconds())
		r = r + 1
	}
}
