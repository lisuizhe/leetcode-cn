package leetcode

/*
 * @lc app=leetcode.cn id=11 lang=golang
 *
 * [11] 盛最多水的容器
 */

// @lc code=start
func maxArea(height []int) int {
	left, right, res := 0, len(height)-1, 0
	for left < right {
		area := min(height[left], height[right]) * (right - left)
		res = max(res, area)
		// moving higher pointer only decrease area
		// so we move lower pointer
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return res
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

// @lc code=end
