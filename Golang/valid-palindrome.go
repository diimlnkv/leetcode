// Time:
// Space:
package main

import "unicode"

func isPalindrome(s string) bool {
	runes := getRunes(s)

	if len(runes) == 0 {
		return true
	}

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		if runes[i] != runes[j] {
			return false
		}

	}
	return true
}

func getRunes(s string) []rune {
	var runes []rune

	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			runes = append(runes, unicode.ToLower(r))
		}
	}
	return runes
}
