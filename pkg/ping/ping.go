package ping

import (
	"fmt"
	"net"
	"time"
)

func TCPPing(destination Dest, pingLimit int, network string) {
	fmt.Printf("PING %s:%s (%s):\n", destination.Hostname, destination.Port, destination.IPv4Addr) //TODO: pick first IPV4 address
	for i := 1; i < pingLimit; i++ {
		time.Sleep(1 * time.Second)
		now := time.Now()
		connDial, err := net.DialTimeout(network, destination.Hostname+":"+destination.Port, 5*time.Second)
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
	}
}
