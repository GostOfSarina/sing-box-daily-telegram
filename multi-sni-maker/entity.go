package main

type RealityJson struct {
	DNS struct {
		Servers  []Server `json:"servers"`
		Strategy string   `json:"strategy"`
	} `json:"dns"`
	Inbounds  []Inbound `json:"inbounds"`
	Outbounds []struct {
		Type string `json:"type"`
		Tag  string `json:"tag,omitempty"`
	} `json:"outbounds"`
	Route struct {
		Rules []Rule `json:"rules"`
	} `json:"route"`
}

type Server struct {
	Address string `json:"address"`
	Detour  string `json:"detour"`
}

type Rule struct {
	Geoip    []string `json:"geoip,omitempty"`
	Outbound string   `json:"outbound"`
	Geosite  []string `json:"geosite,omitempty"`
	IPCidr   []string `json:"ip_cidr,omitempty"`
	Network  string   `json:"network,omitempty"`
	Port     []int    `json:"port,omitempty"`
	Domain   []string `json:"domain,omitempty"`
}

type Inbound struct {
	Type                     string `json:"type"`
	Tag                      string `json:"tag"`
	Listen                   string `json:"listen"`
	ListenPort               int    `json:"listen_port"`
	Sniff                    bool   `json:"sniff"`
	SniffOverrideDestination bool   `json:"sniff_override_destination"`
	DomainStrategy           string `json:"domain_strategy"`
	Users                    []User `json:"users"`
	TLS                      struct {
		Enabled    bool   `json:"enabled"`
		ServerName string `json:"server_name"`
		Reality    struct {
			Enabled   bool `json:"enabled"`
			Handshake struct {
				Server     string `json:"server"`
				ServerPort int    `json:"server_port"`
			} `json:"handshake"`
			PrivateKey string   `json:"private_key"`
			ShortID    []string `json:"short_id"`
		} `json:"reality"`
	} `json:"tls"`
}

type User struct {
	NAME string `json:"name"`
	UUID string `json:"uuid"`
	Flow string `json:"flow"`
}
