package auth

import "regexp"

// WhitelistEntry is a URL matcher that is used to whitelist
// certain HTTP method/URLs. It comprises of a Method string
// (eg, GET) and a regexp for matching the URL.
type WhitelistEntry struct {
	Method string
	URL    *regexp.Regexp
}

// Whitelist is an array of WhitelistEntry's for URLs that require
// no authentication or authorization.
var Whitelist = []WhitelistEntry{

	// User creation/login
	{Method: "PUT", URL: regexp.MustCompile(`/v[0-9]+/users[/]?`)},
	{Method: "POST", URL: regexp.MustCompile(`/v[0-9]+/users[/]?`)},

	// Status (ping)
	{Method: "GET", URL: regexp.MustCompile(`/v[0-9]+/_ping?`)},
}
