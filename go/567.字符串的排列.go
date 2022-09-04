package leetcode

/*
 * @lc app=leetcode.cn id=567 lang=golang
 *
 * [567] 字符串的排列
 */

// @lc code=start
func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	need := map[byte]int{}
	for i := 0; i < len(s1); i++ {
		if count, ok := need[s1[i]-'a']; ok {
			need[s1[i]-'a'] = count + 1
		} else {
			need[s1[i]-'a'] = 1
		}
	}
	window := map[byte]int{}

	left, right, valid := 0, 0, 0
	for right < len(s2) {
		c := s2[right] - 'a'
		right++
		if count, ok := need[c]; ok {
			if wcount, wok := window[c]; wok {
				window[c] = wcount + 1
			} else {
				window[c] = 1
			}
			if count == window[c] {
				valid++
			}
		}

		for right-left >= len(s1) {
			if valid == len(need) {
				return true
			}
			removec := s2[left] - 'a'
			left++
			if removeCount, removeOk := need[removec]; removeOk {
				if removeWcount, removeWok := window[removec]; removeWok {
					if removeCount == removeWcount {
						valid--
					}
					window[removec] = removeWcount - 1
				}
			}
		}
	}

	return false
}

// @lc code=end
