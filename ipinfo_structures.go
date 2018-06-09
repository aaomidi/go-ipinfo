package ipinfo

import "net"

// The full response structure
type IPResponse struct {
	Ip       net.IP  `json:"ip"`
	Hostname string  `json:"hostname"`
	City     string  `json:"city"`
	Region   string  `json:"region"`
	Country  string  `json:"country"`
	Loc      string  `json:"loc"`
	Postal   string  `json:"postal"`
	Org      string  `json:"org"`
	Phone    string  `json:"phone"`
	Asn      Asn     `json:"asn"`
	Company  Company `json:"company"`
	Carrier  Carrier `json:"carrier"`
}

// The ASN structure
type Asn struct {
	Asn    string `json:"asn"`
	Name   string `json:"name"`
	Domain string `json:"domain"`
	Route  string `json:"route"`
	Type   string `json:"type"`
}

// The company structure
type Company struct {
	Name   string `json:"name"`
	Domain string `json:"domain"`
	Type   string `json:"type"`
}

// The carrier structure
type Carrier struct {
	Name string `json:"name"`
	Mcc  string `json:"mcc"`
	Mnc  string `json:"mnc"`
}
