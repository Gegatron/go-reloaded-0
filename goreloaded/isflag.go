package goreloaded

import (
	"strings"
	"unicode"
)

func IsMultiFlag(s string) bool {
	if strings.HasPrefix(s, "(up,") {
		if s[4] != ' ' {
			return false
		}
		if !(s[5] == '-' || s[5] == '+' || unicode.IsNumber(rune(s[5]))) {
			return false
		}
		for i := 6; i < len(s)-1; i++ {
			if !unicode.IsNumber(rune(s[i])) {
				return false
			}
		}
		if s[len(s)-1]!=')' {
			return false
		}
		return true
	}
	if strings.HasPrefix(s, "(low,") || strings.HasPrefix(s, "(cap,") {
		if s[5] != ' ' {
			return false
		}
		if !(s[6] == '-' || s[6] == '+' || unicode.IsNumber(rune(s[6]))) {
			return false
		}
		for i := 7; i < len(s)-1; i++ {
			if !unicode.IsNumber(rune(s[i])) {
				return false
			}
		}
		if s[len(s)-1]!=')' {
			return false
		}
		return true
	}
	return false
}
