package goreloaded

import (
	"strings"
)

func Quotes(reloaded []string) []string {
	s := strings.Join(reloaded, " ")
	var open bool
	var index int
	var str string
	for i, c := range s {
		if i!=len(s)-1 {
			if !open  {
			if c == '\'' &&  (s[i+1] == ' ' || s[i-1] == ' ' || s[i+1] == '\'' || s[i-1] == '\''){
				open = true
				index = i
			} else {
				str += string(c)
			}
		} else if open && c == '\'' {
			str += " " + "'" + s[index:i] + "'" + " "
			open = false
			continue
		}
		}
		
		if i == len(s)-1 && open {
			str += s[index:]
		}

	}

	return strings.Fields(str)
}
