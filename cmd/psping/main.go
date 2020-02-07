package main

import (
    "fmt"
    "net"
    "os"
    "time"
)

var (
    address string
    network = "tcp"
)


func main() {
    i := 0
    targetIP, err := net.LookupHost(target)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Could not get IPs: %v\n", err)
        os.Exit(1)
    }
    fmt.Printf("PING %s (%s) on port %s\n", target, targetIP[0], port) //TODO: pick first IPV4 address
    for {
        time.Sleep(1 * time.Second)
        now := time.Now()
        connDial, err := net.DialTimeout(network, address, 5 * time.Second)
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
        i = i + 1
    }
}
