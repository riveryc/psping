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

func exitReport(hostname, port, network string, total, received *int, min, max, avg *int64) {
	// capture ctrl+c and stop CPU profiler
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for sig := range c {
			fmt.Printf("--- %s:%s %s ping statistics ---\n", hostname, port, network)
			failureRate := (*total - *received) * 100 / *total
			missing := *total - *received
			fmt.Printf("%d packets transmitted, %d packets received, %d packets missing, %d %% packet loss\n", *total, *received, missing, failureRate)
			fmt.Printf("latency Min = %dms, Max = %dms, Average = %dms\n", *min, *max, *avg)
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
	var latencyMin int64 = 9999 // assign a large value that real latency must smaller than this.
	var latencyMax int64 = 0
	var latencyAvg int64 = 0
	exitReport(hostname, port, network, &i, &r, &latencyMin, &latencyMax, &latencyAvg)
	for {
		time.Sleep(1 * time.Second)
		now := time.Now()
		connDial, err := net.DialTimeout(network, hostname+":"+port, 5*time.Second)
		i++
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
		if latency.Milliseconds() > latencyMax {
			latencyMax = latency.Milliseconds()
		}
		if latency.Milliseconds() < latencyMin {
			latencyMin = latency.Milliseconds()
		}
		latencyAvg = (int64(r) * latencyAvg + latency.Milliseconds())/ (int64(r)+1)
		r = r + 1
		connDial.Close()
	}
}
