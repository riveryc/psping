package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

type TargetInput string

type Dest struct {
	Hostname string `json:"hostName"`
	Port     string `json:"port"`
	HasPort  bool   `json:"hasPort"`
	IPv4Addr string `json:"IPv4Addr"`
}

func (i *TargetInput) NewDest() (Dest Dest) {
	hasPort := strings.Contains(string(*i), ":")
	hostname := string(*i)
	port := "80"
	if hasPort {
		index := strings.Index(string(*i), ":")
		hostname = string(*i)[0:index]
		port = string(*i)[index+1:]
	}
	targetIP, err := net.LookupHost(Dest.Hostname)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not get IPs: %v\n", err)
		os.Exit(1)
	}
	Dest.Hostname = hostname
	Dest.HasPort = hasPort
	Dest.Port = port
	Dest.IPv4Addr = targetIP[0]
	return
}
