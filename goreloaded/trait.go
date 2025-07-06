package goreloaded

import (
	"fmt"
	"strings"
)

func Trait(s string) string {

	fixed := clean(strings.Split(s, " "))
	fmt.Println(fixed)

	fmt.Println(len(fixed))
	fixed = FixFlags(fixed)

	fixed = Punc(fixed)
	fixed = ATooAn(fixed)
	fixed = Quotes(fixed)
	fmt.Println(fixed)

	return strings.Join(fixed, " ")
}
func clean(clea []string) []string {
	filtred := []string{}
	for _, str := range clea {
		
		if str != "" {
			filtred = append(filtred, str)
		}
	
	}
	return filtred
}
