package problem1119

import "strings"

func removeVowels(S string) string {
	r := []rune(S)
	vowels := "aeiou"
	ans := make([]rune, len(S))
	idx := 0
	for _, c := range r {
		if !strings.ContainsRune(vowels, c) {
			ans[idx] = c
			idx++
		}
	}
	return string(ans[:idx])
}
