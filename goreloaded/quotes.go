package goreloaded

import (
	"strings"
)

func Quotes(reloaded []string) []string {
	s := strings.Join(reloaded, " ")
	var open bool
	index := 0
	var str string
	for i, c := range s {
		if i != len(s)-1 && i != 0 {
			if !open {
				if c == '\'' && (s[i+1] == ' ' || s[i-1] == ' ' || s[i+1] == '\'' || s[i-1] == '\'' || IsPunc(rune(s[i+1])) || IsPunc(rune(s[i-1]))) {
					str += s[index:i]
					open = true
					index = i
				}
			} else if open && c == '\'' && (s[i+1] == ' ' || s[i-1] == ' ' || s[i+1] == '\'' || s[i-1] == '\'' || IsPunc(rune(s[i+1])) || IsPunc(rune(s[i-1]))) {
				if IsPunc(rune(s[i+1])) || IsPunc(rune(s[i-1])) {
					str += " " + "'" + strings.TrimSpace(s[index+1:i]) + "'"
					open = false
					index = i + 1
					continue
				}
				str += " " + "'" + strings.TrimSpace(s[index+1:i]) + "'" + " "
				open = false
				index = i + 1
				continue
			}
		} else if i == 0 {
			if c == '\'' {
				str += s[index:i]
				open = true
				index = i
			}
		}
		if i == len(s)-1 {
			if open && c == '\'' {
				str += " " + "'" + strings.TrimSpace(s[index+1:len(s)-1]) + "'"
			} else {
				str += s[index:]
			}
		}
	}
	return strings.Fields(str)
}
