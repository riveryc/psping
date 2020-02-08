package ping

type Dest struct {
	Hostname string `json:"hostName"`
	Port     string `json:"port"`
	HasPort  bool   `json:"hasPort"`
	IPv4Addr string `json:"IPv4Addr"`
}
