package solution

import (
	"strconv"
)

func freqAlphabets(s string) string {
	sb := make([]rune, len(s))

	j := 0

	for i := len(s) - 1; i >= 0; i-- {
		char := s[i]
		if char == '#' {
			n := s[i-2 : i]
			num, _ := strconv.Atoi(n)
			sb[j] = rune('a' + uint8(num) - 1)
			i = i - 2
			j++
			continue
		}
		sb[j] = rune('a' + char - '1')
		j++
	}

	sb = sb[0:j]

	for i := len(sb)/2 - 1; i >= 0; i-- {
		index := len(sb) - 1 - i

		sb[i], sb[index] = sb[index], sb[i]
	}

	return string(sb)
}

