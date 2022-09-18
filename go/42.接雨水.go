package leetcode

/*
 * @lc app=leetcode.cn id=42 lang=golang
 *
 * [42] 接雨水
 */

// @lc code=start
func trap(height []int) int {
	// maxLeft is height of max(height[0:left+1]...)
	// maxRight is height of max(height[right:]...)
	left, right, res, maxLeft, maxRight := 0, len(height)-1, 0, 0, 0
	for left < right {
		maxLeft = max(maxLeft, height[left])
		maxRight = max(maxRight, height[right])
		// trapped water in height[i] depends on min(maxLeft. maxRight)
		// so moving the higher side pointer cannot determine next trapped water height
		// we need to move the lower side pointer
		if maxLeft < maxRight {
			res += maxLeft - height[left]
			left++
		} else {
			res += maxRight - height[right]
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

// @lc code=end
