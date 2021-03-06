package mailchecker

import (
	"strings"
)

func Blacklist() []string {
	return blacklist[:]
}

func IsValid(email string) bool {
	return validEmail(email) && !isBlacklisted(email)
}

func validEmail(email string) bool {
	return validMatcherRegex.MatchString(email)
}

func isBlacklisted(email string) bool {

	parts := strings.Split(email, "@")
	domain := parts[1]

	for _, domainSuffix := range allDomainSuffixes(domain) {
		if findBlacklistDomain(domainSuffix) {
			return true
		}
	}
	return false
}

func allDomainSuffixes(domain string) []string {
	components := strings.Split(domain, ".")

	sufixes := make([]string, len(components))

	for i := 0; i < len(components); i++ {
		sufixes = append(sufixes, strings.Join(components[i:], "."))
	}

	return sufixes
}

func findBlacklistDomain(domain string) bool {
	for _, item := range blacklist {
		if item == domain {
			return true
		}
	}
	return false
}
