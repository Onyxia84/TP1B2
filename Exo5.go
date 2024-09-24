package main

func Ft_max_substring(s string) int {
	seen := make(map[byte]int)
	maxLength := 0
	start := 0

	for i := 0; i < len(s); i++ {
		if idx, ok := seen[s[i]]; ok && idx >= start {
			start = idx + 1
		}
		seen[s[i]] = i
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
	}

	return maxLength
}
