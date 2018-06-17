package ipinfo

import (
	"net/http"
	"regexp"
)

var (
	rx = regexp.MustCompile("(?i) bot|spider")
)

// IsBot checks a request's headers to see if it from a crawler.
func IsBot(r *http.Request) bool {

	ua := r.Header.Get("user-agent")

	if rx.FindString(ua) == "" {
		return false
	}

	return true
}
