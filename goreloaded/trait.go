package goreloaded

import (
	"fmt"
	"strings"
)

func Trait(s string) string {
	fixed := strings.Fields(s)
	fmt.Println(fixed)

	fmt.Println(len(fixed))
	fixed = FixFlags(fixed)

	fixed = Punc(fixed)
	fixed = ATooAn(fixed)
	fixed = Quotes(fixed)
	fmt.Println(fixed)

	return strings.Join(fixed, " ")
}

func Traited(s string) string {
	str := ""
	var new string
	for i, c := range s {
		if c != '\n' {
			new += string(c)
		}
		if c == '\n' {
			str += Trait(new) + "\n"
			new = ""
			continue
		}
		if c != '\n' && i == len(s)-1 {
			str += Trait(new)
		}
	}
	return str
}
