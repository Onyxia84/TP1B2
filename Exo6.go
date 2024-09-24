package main

import (
	"math"
)

func Ft_min_window(s string, t string) string {
	if len(t) == 0 || len(s) < len(t) {
		return ""
	}

	countT := make(map[rune]int)
	for _, char := range t {
		countT[char]++
	}

	countWindow := make(map[rune]int)
	formed := 0
	required := len(countT)

	l, r := 0, 0
	minLen := math.MaxInt32
	minLeft := 0

	for r < len(s) {
		char := rune(s[r])
		countWindow[char]++

		if countT[char] > 0 && countWindow[char] == countT[char] {
			formed++
		}

		for l <= r && formed == required {
			char = rune(s[l])

			if r-l+1 < minLen {
				minLen = r - l + 1
				minLeft = l
			}

			countWindow[char]--
			if countT[char] > 0 && countWindow[char] < countT[char] {
				formed--
			}
			l++
		}
		r++
	}

	if minLen == math.MaxInt32 {
		return ""
	}
	return s[minLeft : minLeft+minLen]
}
