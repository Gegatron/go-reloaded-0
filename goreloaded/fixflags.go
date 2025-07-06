package goreloaded

import (
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
			n, err := GetNumber(c[i+1])
			if err != nil {
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

func GetNumber(s string) (int, error) {
	new := ""
	b := false
	for i := 0; i < len(s); i++ {
		if !b && (s[i] == '-' || (s[i] >= '0' && s[i] <= '9')) && new == "" {
			new += string(s[i])
			b = true
			continue
		}
		if b && i != len(s)-1 {
			new += string(s[i])
		}
	}
	return strconv.Atoi(new)
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
		return s
	}
	return strconv.Itoa(int(hex))
}

func ToBin(s string) string {
	bin, err := strconv.ParseInt(s, 2, 64)
	if err != nil {
		return s
	}
	return strconv.Itoa(int(bin))
}
