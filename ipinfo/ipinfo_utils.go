package ipinfo

import (
	"net/http"
	"regexp"
)

var (
	rx = regexp.MustCompile("(?i) bot|spider")
)

func IsBot(r *http.Request) bool {

	ua := r.Header.Get("user-agent")

	if rx.FindString(ua) == "" {
		return false
	}

	return true
}
