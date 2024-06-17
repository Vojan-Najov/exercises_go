package main

import (
	"unicode"
)

func DeleteVowels(s string) string {
	res := make([]rune, 0, len(s))
	for _, r := range s {
		lc := unicode.ToLower(r)
		if lc == 'a' || lc == 'e' || lc == 'i' || lc == 'o' || lc == 'u' {
			continue
		}
		res = append(res, r)
	}
	return string(res)
}
