package leetcode

/*
 * @lc app=leetcode.cn id=27 lang=golang
 *
 * [27] 移除元素
 */

// @lc code=start
func removeElement(nums []int, val int) int {
	n := len(nums)
	count := 0
	for i := 0; i < n; i++ {
		if nums[i] != val {
			count++
			nums[count-1] = nums[i]
		}
	}
	return count
}

// @lc code=end
