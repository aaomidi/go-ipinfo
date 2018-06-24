package ipinfo

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/patrickmn/go-cache"
)

const (
	// BaseURL is the base URL of the api endpoint.
	BaseURL = "https://ipinfo.io/%s"
)

// IPInfo is the entry point to the API wrapper.
type IPInfo struct {
	Token          string
	Client         *http.Client
	LanguageReader *io.Reader

	Cache *cache.Cache

	Lang       map[string]string
	langFailed bool
}

// init initializes the http client if it doesn't exist.
func (i *IPInfo) init() {
	// Sets up the client
	if i.Client == nil {
		i.Client = http.DefaultClient
		i.Client.Timeout = 5 * time.Second
	}

	if i.Cache == nil {
		i.Cache = cache.New(1*time.Minute, 10*time.Minute)
	}

	// Sets up the language reader
	if i.LanguageReader == nil {
		var file io.Reader
		file, err := os.Open("../static/en_US.json")
		if err != nil {
			fmt.Println(err)
			i.langFailed = true
		}

		i.LanguageReader = &file
	}

	// Sets up the conversion map
	if i.Lang == nil && i.LanguageReader != nil && !i.langFailed {
		i.Lang = make(map[string]string)

		langFile, err := ioutil.ReadAll(*i.LanguageReader)
		if err != nil {
			fmt.Println(err)
			i.langFailed = true
			return
		}

		err = decode(langFile, &i.Lang)
		if err != nil {
			fmt.Println(err)
			i.langFailed = true
		}
	}
}

// LookupASN calls the ASN lookup feature of the wrapper
func (i *IPInfo) LookupASN(asn string) (*ASNResponse, error) {
	i.init()

	if resp, ok := i.Cache.Get(fmt.Sprintf("ASN-%s", asn)); ok {
		result := resp.(*ASNResponse)
		return result, nil
	}

	url := ""
	if asn == "" {
		return nil, fmt.Errorf("asn can not be null")
	}

	url = fmt.Sprintf(BaseURL, fmt.Sprintf("%s/json", asn))
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	i.prepareRequest(req)

	resp, err := i.Client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusTooManyRequests {
			return nil, NewRateLimitedError()
		}
		return nil, fmt.Errorf("received statuscode: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	errResponse := ErrorResponse{}
	err = decode(body, &errResponse)

	if err == nil && errResponse.Error != "" {
		return nil, NewErrorResponseError(&errResponse)
	}

	response := ASNResponse{}
	err = decode(body, &response)
	if err != nil {
		return nil, err
	}

	i.Cache.Set(fmt.Sprintf("ASN-%s", asn), &response, cache.DefaultExpiration)

	return &response, nil
}

// LookupIP looks up the IPResponse from an IP.
func (i *IPInfo) LookupIP(ip net.IP) (*IPResponse, error) {
	i.init()

	if resp, ok := i.Cache.Get(fmt.Sprintf("IP-%s", ip.String())); ok {
		result := resp.(*IPResponse)
		return result, nil
	}

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

	i.prepareRequest(req)

	resp, err := i.Client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusTooManyRequests {
			return nil, NewRateLimitedError()
		}
		return nil, fmt.Errorf("received statuscode: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	errResponse := ErrorResponse{}
	err = decode(body, &errResponse)

	if err == nil && errResponse.Error != "" {
		return nil, NewErrorResponseError(&errResponse)
	}

	response := IPResponse{}
	err = decode(body, &response)
	if err != nil {
		return nil, err
	}

	i.Cache.Set(fmt.Sprintf("IP-%s", ip.String()), &response, cache.DefaultExpiration)

	return &response, nil
}

// GetCountry gets the full name of the country out of the ISO2 CountryCode
func (i *IPInfo) GetCountry(code string) (string, error) {
	i.init()
	if i.langFailed {
		return "", fmt.Errorf("country code file failed to load - check logs")
	}

	data, ok := i.Lang[code]
	if !ok {
		return "", NewNoSuchCountryError(code)
	}
	return data, nil
}

// decode decodes the json response from the server.
func decode(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

// prepareRequest prepares the request to be sent to IPInfo
func (i *IPInfo) prepareRequest(req *http.Request) {
	req.SetBasicAuth(i.Token, "")
	req.Header.Set("user-agent", "IPInfoClient/GoLang/0.1")
}
