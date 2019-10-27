package medium

// LengthOfLongestSubstring is 滑动区间 每次取当前截取字符的最后一位，如果存在则start +1,且剔除map 当前第一位字符 不存在则end +1,
func LengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	start, end, maxLen := 0, 1, 0
	map1 := make(map[string]int)
	for start <= len(s) && end <= len(s) {
		char := s[start:end]
		lastChar := char[len(char)-1 : len(char)]
		if map1[lastChar] > 0 {
			map1[s[start:start+1]] = 0
			start++
		} else {
			map1[lastChar] = 1
			end++
			maxLen = max(maxLen, len(char))
		}

	}
	return maxLen
}

// 别人的解法，不知道为什么是128？应该是256个字符吧，和我的解法一个思路，不过算法实现更优化
func lengthOfLongestSubstring(s string) int {
	maxLen, start := 0, 0
	table := [128]int{}
	for i := range table {
		table[i] = -1
	}
	for i, c := range s {
		if table[c] >= start {
			start = table[c] + 1
		}
		table[c] = i
		maxLen = max(maxLen, i-start+1)
	}
	return maxLen
}
func max(x int, y int) int {
	if x > y {
		return x
	}
	return y

}
