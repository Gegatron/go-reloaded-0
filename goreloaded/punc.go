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
		if i != 0 && IsPunc(rune(strs[i][0])) {
			strs[i-1] = strs[i-1] + strs[i]
			strs[i] = ""
			strs = strings.Fields(strings.Join(strs, " "))
			i--
		}
	}
	return strs
}

func IsPunc(s rune) bool {
	return s == ',' || s == '.' || s == ':' || s == ';' || s == '?' || s == '!'
}
