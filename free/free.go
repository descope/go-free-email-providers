// Package free provides functions to check if the domain of an email or a domain is in the community maintained https://f.hubspotusercontent40.net/hubfs/2832391/Marketing/Lead-Capture/free-domains-2.csv free list.
package free

import (
	"strings"
)

// IsFreeEmail checks if the domain of an email address is in a list of free email domain names
func IsFreeEmail(email string) bool {
	at := strings.LastIndex(email, "@")
	if at == -1 {
		return true // not a valid email,
	}
	domain := email[at+1:]

	_, ok := domains[domain]
	return ok
}

// IsBurnerDomain checks if a domain is in a list of free email domain names
func IsFreeDomain(domain string) bool {
	_, ok := domains[domain]
	return ok
}
