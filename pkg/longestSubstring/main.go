package longestsubstring

var s = "abcabcbb"

func lengthOfLongestSubstring(s string) int {
	var ans int
	var m = make(map[byte]int)
	for i, j := 0, 0; j < len(s); j++ {
		if _, ok := m[s[j]]; ok {
			i = max(i, m[s[j]])
		}
		ans = max(ans, j-i+1)
		m[s[j]] = j + 1
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Ans() {
	println(lengthOfLongestSubstring(s))
}
