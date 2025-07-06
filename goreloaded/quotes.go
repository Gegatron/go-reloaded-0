package goreloaded

import (
	"strings"
)

func Quotes(reloaded []string) []string {
	s := strings.Join(reloaded, " ")
	index := 0
	str := ""
	b := true
	for i, c := range s {
		if !b && c == '\n' {
			str += strings.TrimSpace(s[index:i])
			b = true
		}
		if b && c == '\'' {
			if i == len(s)-1 {
				str += string(c)
				continue
			}
			if i == 0 {
				str += " "
				index = i
				b = false
				continue
			}
			if b && s[i+1] == ' ' || s[i-1] == ' ' || (s[i+1] == '\'' || s[i-1] == '\'') {
				for j := i + 1; j < len(s); j++ {
					if s[j] == '\n' {
						break
					}
					if s[j] == '\'' {
						str += " "
					}
				}
				index = i
				b = false
				continue
			}
		}
		if !b && c == '\'' {
			if i != len(s)-1 && (s[i+1] == ' ' || s[i-1] == ' ' || (s[i+1] == '\'' || s[i-1] == '\'')) {
				str += "'" + strings.TrimSpace(s[index+1:i]) + "'" + " "
				b = true

			} else if i == len(s)-1 {
				str += "'" + strings.TrimSpace(s[index+1:i]) + "'"
			}
			continue
		} else if b {
			str += string(c)
		} else if !b && i == len(s)-1 {
			str += s[index:]
		}
	}
	reloaded = strings.Split(str," ")
	return reloaded
}
