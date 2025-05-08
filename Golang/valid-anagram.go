// Time:
// Space:
package main

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	runes := map[rune]int{}

	for _, r := range s {
		runes[r]++
	}
	for _, r := range t {
		if c, ok := runes[r]; !ok || c == 0 {
			return false
		}
		runes[r]--
	}
	return true
}
