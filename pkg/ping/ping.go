package ping

import (
	"fmt"
	"net"
	"time"
)

func TCPPing(hostname, port, IPv4Addr, network string, pingLimit int) {
	fmt.Printf("PING %s:%s (%s):\n", hostname, port, IPv4Addr) //TODO: pick first IPV4 address
	for i := 1; i < pingLimit; i++ {
		time.Sleep(1 * time.Second)
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
	}
}
