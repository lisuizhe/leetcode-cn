package leetcode

/*
 * @lc app=leetcode.cn id=76 lang=golang
 *
 * [76] 最小覆盖子串
 */

// @lc code=start
func minWindow(s string, t string) string {
	if len(s) == 0 || len(t) == 0 || len(s) < len(t) {
		return ""
	}
	need := map[byte]int{}
	for i := 0; i < len(t); i++ {
		if count, ok := need[t[i]-'a']; ok {
			need[t[i]-'a'] = count + 1
		} else {
			need[t[i]-'a'] = 1
		}
	}
	window := map[byte]int{}

	left, right, valid, start, resLen := 0, 0, 0, 0, len(s)+1
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

		for valid == len(need) {
			if right-left < resLen {
				start = left
				resLen = right - left
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

	if resLen == len(s)+1 {
		return ""
	} else {
		return s[start:(start + resLen)]
	}
}

// @lc code=end
