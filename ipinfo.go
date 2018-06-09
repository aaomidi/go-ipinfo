package ipinfo

import (
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"net/http"
)

const (
	BaseURL = "https://ipinfo.io/%s"
)

// The IPInfo initializer
type IPInfo struct {
	Token  string
	Client *http.Client
}

// init initializes the http client if it doesn't exist
func (i *IPInfo) init() {
	if i.Client == nil {
		i.Client = http.DefaultClient
	}
}

func (i *IPInfo) LookupIP(ip *net.IP) (IPResponse, error) {
	i.init()
	var response IPResponse

	url := ""

	if ip == nil {
		url = fmt.Sprintf(BaseURL, "json")
	} else {
		url = fmt.Sprintf(BaseURL, fmt.Sprintf("%s/json", ip.String()))
	}

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return response, err
	}
	req.SetBasicAuth(i.Token, "")

	fmt.Println(req)
	resp, err := i.Client.Do(req)

	if err != nil {
		return response, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return response,
			errors.New(fmt.Sprintf("Received statuscode: %d", resp.StatusCode))
	}

	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return response, nil
	}

	return response, nil
}
