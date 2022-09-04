package leetcode

/*
 * @lc app=leetcode.cn id=438 lang=golang
 *
 * [438] 找到字符串中所有字母异位词
 */

// @lc code=start
func findAnagrams(s string, p string) []int {
	if len(s) < len(p) || len(s) == 0 || len(p) == 0 {
		return []int{}
	}

	need := map[byte]int{}
	for i := 0; i < len(p); i++ {
		if count, ok := need[p[i]-'a']; ok {
			need[p[i]-'a'] = count + 1
		} else {
			need[p[i]-'a'] = 1
		}
	}
	window := map[byte]int{}

	left, right, valid, res := 0, 0, 0, []int{}
	for right < len(s) {
		c := s[right] - 'a'
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

		for right-left >= len(p) {
			if valid == len(need) {
				res = append(res, left)
			}

			removec := s[left] - 'a'
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
	return res
}

// @lc code=end
