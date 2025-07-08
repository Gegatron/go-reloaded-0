package goreloaded

import (
	"strings"
)

func Trait(s string) string {
	fixed := strings.Fields(s)
	fixed = FixFlags(fixed)
	fixed = Punc(fixed)
	fixed = Quotes(fixed)

	fixed = ATooAn(fixed)
	return strings.Join(fixed, " ")
}

func Traited(s string) string {
	str := ""
	Traited := strings.Split(s, "\n")
	for i, c := range Traited {
		if i != len(Traited)-1 {
			str += Trait(c) + "\n"
			continue
		}
		str += Trait(c)
	}
	return str
}
