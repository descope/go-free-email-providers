package free

import (
	"testing"
)

func TestNonFreeEmails(t *testing.T) {
	nonBurnerEmails := []string{
		"meir@descope.com",
		"1@corp1.com",
	}
	for _, email := range nonBurnerEmails {
		if ok := IsFreeEmail(email); ok {
			t.Errorf("%s should not be considered burner email", email)
		}
	}
}

func TestFreeEmails(t *testing.T) {
	burnerEmails := []string{
		"hello@protonmail.com",
		"hello@gmail.com",
		"hello@yahoo.com",
		"hello@gmail.com",
		"my.name@hotmail.com",
		"hello@outlook.com",
		"hello@protonmail.com",
	}
	for _, email := range burnerEmails {
		if ok := IsFreeEmail(email); !ok {
			t.Errorf("%s should be considered burner email", email)
		}
	}
}

func TestNonFreeDomains(t *testing.T) {
	nonBurnerDomains := []string{
		"descope.com",
		"corp1.com",
	}
	for _, email := range nonBurnerDomains {
		if ok := IsFreeDomain(email); ok {
			t.Errorf("%s should not be considered burner domain", email)
		}
	}
}

func TestFreeDomains(t *testing.T) {
	burnerEmails := []string{
		"protonmail.com",
		"gmail.com",
	}
	for _, email := range burnerEmails {
		if ok := IsFreeDomain(email); !ok {
			t.Errorf("%s should be considered burner domain", email)
		}
	}
}
