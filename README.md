# GO IPInfo
[![License](http://img.shields.io/:license-mit-blue.svg)](LICENSE)
[![Travis](https://travis-ci.com/aaomidi/go-ipinfo.svg?branch=master&style=flat-square)](https://travis-ci.com/aaomidi/go-ipinfo)
[![Go Report Card](https://goreportcard.com/badge/github.com/aaomidi/go-ipinfo)](https://goreportcard.com/report/github.com/aaomidi/go-ipinfo)
[![GoDoc](https://godoc.org/github.com/aaomidi/go-ipinfo?status.svg)](https://godoc.org/github.com/aaomidi/go-ipinfo)


An unofficial GoLang wrapper around IPInfo.

## Features:
- IP Lookup
- ASN Lookup

## Installing it

````bash
go get github.com/aaomidi/go-ipinfo/...
````

## Using it

To get information about an IP:

````golang
func main() {
    api := IPInfo{Token: "YOUR_TOKEN"}

    ip := net.ParseIP("8.8.8.8")

    response, err := api.LookupIP(ip) // Type IPResponse
}
````
