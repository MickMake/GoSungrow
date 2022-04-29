package api

import (
	"strings"
)


func CleanString(s string) string {
	// var ret string
	var result strings.Builder
	for i := 0; i < len(s); i++ {
		b := s[i]
		if ('a' <= b && b <= 'z') ||
			('A' <= b && b <= 'Z') ||
			('0' <= b && b <= '9') ||
			(b == '-') ||
			(b == '_') ||
			(b == '.') ||
			b == ' ' {
			result.WriteByte(b)
		}
	}

	// dupes := regexp.MustCompile(`\s+`)
	// ret = dupes.ReplaceAllString(result.String(), )
	return result.String()
}
