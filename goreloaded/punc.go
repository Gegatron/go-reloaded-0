package goreloaded

import (
	"strings"
)

func Punc(strs []string) []string {
	str := strings.Join(strs, " ")
	s := ""
	for i, c := range str {
		if i != len(str)-1 && IsPunc(c) && str[i+1] != ' ' {
			s += string(c) + " "
		} else {
			s += string(c)
		}
	}
	strs = strings.Split(s, " ")
	for i := 0; i < len(strs); i++ {
		if i != 0 && AllPunc(strs[i]) {
			if strs[i-1][len(strs[i-1])-1] == '\n' {
				continue
			}
			strs[i-1] = strs[i-1] + strs[i]
			strs[i] = ""
			strs = clean(strs)
			i--
		}
	}
	return strs
}
func IsPunc(s rune) bool {
	return s == ',' || s == '.' || s == ':' || s == ';' || s == '?' || s == '!'
}
func AllPunc(s string) bool {
	for _, c := range s {
		if !strings.Contains(".,;:!?", string(c)) {
			return false
		}
	}
	return true
}
