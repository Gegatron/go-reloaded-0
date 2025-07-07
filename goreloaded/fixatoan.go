package goreloaded

import (
	"strings"
)

func ATooAn(s []string) []string {
	for i := 0; i < len(s); i++ {
		if i != len(s)-1 && s[i] == "a" {
			if strings.Contains("aeoiuhAEOIUH", string(s[i+1][0])) {
				s[i] = "an"
			}
		}
		if i != len(s)-1 && s[i] == "A" {
			if strings.Contains("aeoiuhAEOIUH", string(s[i+1][0])) {
				s[i] = "An"
			}
		}
	}
	return s
}
