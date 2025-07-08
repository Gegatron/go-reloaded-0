package goreloaded

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func FixFlags(c []string) []string {
	for i := 0; i < len(c); i++ {
		if c[i] == "(up)" || c[i] == "(low)" || c[i] == "(cap)" || c[i] == "(hex)" || c[i] == "(bin)" {
			if i != 0 {
				if c[i] == "(up)" {
					c[i-1] = strings.ToUpper(c[i-1])
				} else if c[i] == "(low)" {
					c[i-1] = strings.ToLower(c[i-1])
				} else if c[i] == "(cap)" {
					c[i-1] = Capitalize(c[i-1])
				} else if c[i] == "(bin)" {
					c[i-1] = ToBin(c[i-1])
				} else if c[i] == "(hex)" {
					c[i-1] = ToHex(c[i-1])
				}
			}
			c[i] = ""
			c = strings.Fields(strings.Join(c, " "))
			i--
			continue
		}
		if i == len(c)-1 {
			continue
		}
		if IsMultiFlag(c[i] + " " + c[i+1]) {
			n, err := strconv.Atoi(c[i+1][:len(c[i+1])-1])
			if err != nil {
				fmt.Println("Error:",err)
				continue
			}
			if n < 0 {
				n = 0
			}
			for j := i - 1; j >= 0; j-- {
				if n > 0 {
					if strings.HasPrefix(c[i], "(up,") {
						c[j] = strings.ToUpper(c[j])
					} else if strings.HasPrefix(c[i], "(low,") {
						c[j] = strings.ToLower(c[j])
					} else if strings.HasPrefix(c[i], "(cap,") {
						c[j] = Capitalize(c[j])
					}
					n--
					continue
				} else {
					break
				}
			}
			c[i] = ""
			c[i+1] = ""
			c = strings.Fields(strings.Join(c, " "))
			i--
			continue
		}
	}
	return c
}


func Capitalize(s string) string {
	new := []rune(s)
	b := false
	for i, c := range new {
		if unicode.IsLetter(c) && !b {
			new[i] = unicode.ToUpper(new[i])
			b = true
		} else if unicode.IsLetter(c) && b {
			new[i] = unicode.ToLower(new[i])
		}
	}
	return string(new)
}

func ToHex(s string) string {
	hex, err := strconv.ParseInt(s, 16, 64)
	if err != nil {
		fmt.Println("Error:",err)
		return s
	}
	return strconv.Itoa(int(hex))
}

func ToBin(s string) string {
	bin, err := strconv.ParseInt(s, 2, 64)
	if err != nil {
		fmt.Println("Error:",err)
		return s
	}
	return strconv.Itoa(int(bin))
}

func IsMultiFlag(s string) bool {
	return (strings.HasPrefix(s, "(low,") || strings.HasPrefix(s, "(cap,") || strings.HasPrefix(s, "(up,")) && s[len(s)-1] == ')'
}
