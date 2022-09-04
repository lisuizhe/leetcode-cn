package leetcode

/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 */

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	left, right, res := 0, 0, 0
	window := map[byte]int{}
	for right < len(s) {
		c := s[right] - 'a'
		right++
		if count, ok := window[c]; ok {
			window[c] = count + 1
		} else {
			window[c] = 1
		}

		for window[c] > 1 {
			removec := s[left] - 'a'
			left++
			if removeCount, removeOk := window[removec]; removeOk {
				window[removec] = removeCount - 1
			}
		}

		res = max(res, right-left)
	}
	return res
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

// @lc code=end
