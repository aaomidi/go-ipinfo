package ipinfo

import "net"

// IPResponse is the full response structure.
type IPResponse struct {
	IP       net.IP  `json:"ip"`
	Hostname string  `json:"hostname"`
	City     string  `json:"city"`
	Region   string  `json:"region"`
	Country  string  `json:"country"`
	Loc      string  `json:"loc"`
	Postal   string  `json:"postal"`
	Org      string  `json:"org"`
	Phone    string  `json:"phone"`
	ASN      Asn     `json:"asn"`
	Company  Company `json:"company"`
	Carrier  Carrier `json:"carrier"`
}

// Asn is the ASN structure.
type Asn struct {
	ASN    string `json:"asn"`
	Name   string `json:"name"`
	Domain string `json:"domain"`
	Route  string `json:"route"`
	Type   string `json:"type"`
}

// Company is the company structure.
type Company struct {
	Name   string `json:"name"`
	Domain string `json:"domain"`
	Type   string `json:"type"`
}

// Carrier is the carrier structure.
type Carrier struct {
	Name string `json:"name"`
	MCC  string `json:"mcc"`
	MNC  string `json:"mnc"`
}
