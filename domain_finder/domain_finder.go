package domain_finder

import (
	"strings"
)

func DomainFromUrl(url string) string {
	result := ""
	parts := strings.Split(url, "/")
	if parts[0] == "http:" || parts[0] == "https:" {
		for index := 1; index < len(parts); index++ {
			if parts[index] != "/" && parts[index] != "" {
				result = parts[index]
				break
			}
		}
	}

	parts = strings.Split(result, ".")
	if len(parts) > 2 {
		start := len(parts) - 2
		end := len(parts)
		if len(parts[end-1]) == 2 && len(parts[end-2]) == 2 {
			start--
		}

		result = strings.Join(parts[start:end], ".")
	}

	return result
}
