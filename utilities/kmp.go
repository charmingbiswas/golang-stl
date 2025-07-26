package utilities

func lps(pattern string) []int {
	LPS := make([]int, len(pattern))
	LPS[0] = 0
	i := 1
	length := 0

	for i < len(pattern) {
		if pattern[i] == pattern[length] {
			length++
			LPS[i] = length
			i++
		} else {
			if length != 0 {
				length = LPS[length-1]
			} else {
				LPS[i] = 0
				i++
			}
		}
	}

	return LPS
}

func KmpPatternMatching(input string, pattern string) []int {
	LPS := lps(pattern)
	M := len(input)
	N := len(pattern)
	i := 0
	j := 0
	ans := make([]int, 0)
	for i < M {
		if input[i] == pattern[j] {
			i++
			j++
		}

		if j == N {
			ans = append(ans, i-j)
			j = LPS[j-1]
		} else if input[i] != pattern[j] {
			if j != 0 {
				j = LPS[j-1]
			} else {
				i++
			}
		}
	}
	return ans
}
