// Package netvalidation - validate IP
package netvalidation

import (
	"fmt"
	"net"
	"net/url"
)

// ValidateIP - validate IP
//
// param: host
// return: bool
func ValidateIP(host string) bool {
	return net.ParseIP(host) == nil
}

// CheckDNSLookup - check if DNS the give hosts exists
//
// param: host
// return: list net.IP
// return: error
func CheckDNSLookup(host string) ([]net.IP, error) {
	return net.LookupIP(host)
}

// ValidateURL - check if an URL is valid
//
// param: host
// return: string
// return: error
func ValidateURL(host string) (string, error) {
	u, err := url.ParseRequestURI(host)
	return fmt.Sprint(u), err
}

// ParseURL - parse host in URL
//
// param: host
// return: url
// return: error
func ParseURL(host string) (*url.URL, error) {
	return url.Parse(host)
}
