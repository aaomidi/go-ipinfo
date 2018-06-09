package ipinfo

import (
	"encoding/json"
	"net"
	"net/http"
	"testing"
)

// TestWithClient ensures you can set your own client.
func TestWithClient(t *testing.T) {
	t.Parallel()
	testGoogleDNS(IPInfo{Client: http.DefaultClient}, t)

}

// TestWithoutClient ensures the client is being nullchecked.
func TestWithoutClient(t *testing.T) {
	t.Parallel()
	testGoogleDNS(IPInfo{}, t)
}

// TestBaseURL ensures the BaseURL isn't being changed.
func TestBaseURL(t *testing.T) {
	if BaseURL != "https://ipinfo.io/%s" {
		t.Log("BaseURL was changed without the test being changed.")
		t.Fail()
	}
}

// testGoogleDNS tests to make sure the response returned by the API is valid. This may fail since we're actually doing a call to an external service.
func testGoogleDNS(api IPInfo, t *testing.T) {

	ip := net.ParseIP("8.8.8.8")

	response, err := api.LookupIP(&ip)

	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}

	if response.Hostname != "google-public-dns-a.google.com" {
		t.Log("Google DNS did not return correct hostname.")
		t.Fail()
		return
	}

	if !response.Ip.Equal(ip) {
		t.Log("Google DNS did not have the correct IP.")
		t.Fail()
		return
	}
}

// TestDecoder tests the decoder capability
func TestDecoder(t *testing.T) {
	t.Parallel()

	fake := IPResponse{
		Ip:       net.ParseIP("8.8.8.8"),
		Hostname: "Example",
	}

	var inter IPResponse

	b, _ := json.Marshal(fake)
	decode(b, &inter)

	if !inter.Ip.Equal(fake.Ip) {
		t.Log("IP was not decoded and encoded properly.")
		t.Fail()
		return
	}

	if inter.Hostname != fake.Hostname {
		t.Log("Hostname was not decoded and encoded properly.")
		t.Fail()
		return
	}
}
