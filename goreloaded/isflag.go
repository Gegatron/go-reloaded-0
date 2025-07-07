package goreloaded

import (
	"strings"
)

func IsMultiFlag(s string) bool {

	if strings.HasPrefix(s, "(low,") || strings.HasPrefix(s, "(cap,") || strings.HasPrefix(s, "(up,") {

		if s[len(s)-1] != ')' {
			return false
		}
		return true
	}
	return false
}
