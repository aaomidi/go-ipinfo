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

// TestWithoutClient ensures the client is being null checked.
func TestWithoutClient(t *testing.T) {
	t.Parallel()
	testGoogleDNS(IPInfo{}, t)
}

// TestBaseURL ensures the BaseURL isn't being changed.
func TestBaseURL(t *testing.T) {
	// This test might look ridiculous...
	// But since an update to this can break the entire program it acts as a verification prompt so they don't accidentally break everything.
	if BaseURL != "https://ipinfo.io/%s" {
		t.Error("BaseURL was changed without the test being changed.")
	}
}

// testGoogleDNS tests to make sure the response returned by the API is valid. This may fail since we're actually doing a call to an external service.
func testGoogleDNS(api IPInfo, t *testing.T) {

	ip := net.ParseIP("8.8.8.8")

	response, err := api.LookupIP(ip)

	if err != nil {
		t.Error(err)
		return
	}

	if response.Hostname != "google-public-dns-a.google.com" {
		t.Error("Google DNS did not return correct hostname.")
		return
	}

	if !response.IP.Equal(ip) {
		t.Error("Google DNS did not have the correct IP.")
		return
	}
}

// TestDecoder tests the decoder capability.
func TestDecoder(t *testing.T) {
	t.Parallel()

	fake := IPResponse{
		IP:       net.ParseIP("8.8.8.8"),
		Hostname: "Example",
	}

	var inter IPResponse

	b, _ := json.Marshal(fake)
	decode(b, &inter)

	if !inter.IP.Equal(fake.IP) {
		t.Error("IP was not decoded and encoded properly.")
		return
	}

	if inter.Hostname != fake.Hostname {
		t.Error("Hostname was not decoded and encoded properly.")
		return
	}
}

// TestGetCountry tests the GetCountry method of the API wrapper.
func TestGetCountry(t *testing.T) {
	t.Parallel()
	ip := IPInfo{}

	country, _ := ip.GetCountry("US")

	if country != "United States" {
		t.Error("Country name was invalid.")
	}
}

func TestBotDetection(t *testing.T) {
	t.Parallel()

	req, _ := http.NewRequest("GET", "", nil)
	req.Header.Set("user-agent", "spider")

	if !IsBot(req) {
		t.Error("User-Agent was a bot user-agent./1")
	}

	req.Header.Set("user-agent", "something something BoT something")

	if !IsBot(req) {
		t.Error("User-Agent was a bot user-agent./2")
	}

	req.Header.Set("user-agent", "something something SPiDer something")

	if !IsBot(req) {
		t.Error("User-Agent was a bot user-agent./3")
	}

	req.Header.Set("user-agent", "something something hello something")

	if IsBot(req) {
		t.Error("User-Agent was not bot user-agent./4")
	}

}
