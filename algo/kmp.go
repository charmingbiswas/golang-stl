package algo

func lps(pattern string) []int {
	n := len(pattern)
	i := 1
	length := 0
	LPS := make([]int, n)
	LPS[0] = 0

	for i < n {
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

func KnuthMorrisPrattStringMatching(input string, pattern string) []int {
	ans := make([]int, 0)
	if len(input) < len(pattern) || (len(input) == 0 || len(pattern) == 0) {
		return ans
	}
	LPS := lps(pattern)
	M := len(input)
	N := len(pattern)
	i := 0
	j := 0

	for i < M {
		if input[i] == pattern[j] {
			i++
			j++
		}

		if j == N {
			ans = append(ans, i-j)
			j = LPS[j-1]
		} else if i < M && input[i] != pattern[j] {
			if j != 0 {
				j = LPS[j-1]
			} else {
				i++
			}
		}
	}

	return ans
}
