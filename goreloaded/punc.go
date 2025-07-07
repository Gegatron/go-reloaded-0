package goreloaded

import (
	"strings"
)

func Punc(strs []string) []string {
	str := strings.Join(strs, " ")
	s := ""
	for _, c := range str {
		if IsPunc(c) {
			s += string(c) + " "
		} else {
			s += string(c)
		}
	}
	strs = strings.Fields(s)
	for i := 0; i < len(strs); i++ {
		if i != 0 && AllPunc(strs[i]) {
			strs[i-1] = strs[i-1] + strs[i]
			strs[i] = ""
			strs = strings.Fields(strings.Join(strs," "))
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
