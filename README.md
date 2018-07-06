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

### Examples

To get information about an IP:

````go
func main() {

    api := IPInfo{Token: "YOUR_TOKEN"}

    ip := net.ParseIP("8.8.8.8")

    response, err := api.LookupIP(ip) // Type IPResponse
}
````

To get the information about an ASN:

````
func main() {
    api := IPInfo{Token: "YOUR_TOKEN"}

    response, err := api.LookupASN("AS7922")
}
````

### Errors

We've provided several errors to help you check for common issues:

- RateLimitedError - This error is telling you that you've been limited by the wrapper.
- ErrorResponseError - This error is telling you that the response from one of the API endpoints was an error.
- NoSuchCountryError - This error is telling you that the ISO2 country code couldn't be translated to a country name.

### IPInfo Structure

The struct `IPInfo` is the entry point to this library. None of the values in the struct are required, and will be set to defaults if you do not set them.

Consult the table to see what the default value of these variables are:

| Variable       | Purpose                                                                            | Default Value                                       |
|----------------|------------------------------------------------------------------------------------|-----------------------------------------------------|
| Token          | Token used to communicate with the API                                             | Empty string ""                                     |
| Client         | A pointer to the http.Client.                                                      | The default http.Client with a timeout of 5 seconds |
| LanguageReader | A pointer to io.Reader. This reader will be used to decode the json language file. | A reader to static/en_US.json                       |

### The other structs

[Code](https://github.com/aaomidi/go-ipinfo/blob/master/ipinfo/ipinfo_structures.go) is the best documentation for this.

### The language system

We've included an [example](https://github.com/aaomidi/go-ipinfo/blob/master/static/en_US.json) English, United States language file. You can include your own and provide the reader to us using:
`os.Open("path/to/file")`