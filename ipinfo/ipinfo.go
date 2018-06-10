package ipinfo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
)

const (
	// BaseURL is the base URL of the api endpoint.
	BaseURL = "https://ipinfo.io/%s"
)

// IPInfo is the entry point to the API wrapper.
type IPInfo struct {
	Token  string
	Client *http.Client
}

// init initializes the http client if it doesn't exist.
func (i *IPInfo) init() {
	if i.Client == nil {
		i.Client = http.DefaultClient
	}
}

// LookupIP looks up the IPResponse from an IP.
func (i *IPInfo) LookupIP(ip net.IP) (*IPResponse, error) {
	i.init()

	url := ""

	if ip == nil {
		url = fmt.Sprintf(BaseURL, "json")
	} else {
		url = fmt.Sprintf(BaseURL, fmt.Sprintf("%s/json", ip.String()))
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(i.Token, "")

	resp, err := i.Client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received statuscode: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	response := IPResponse{}
	err = decode(body, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// decode decodes the json response from the server.
func decode(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}
