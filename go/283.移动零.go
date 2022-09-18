package leetcode

/*
 * @lc app=leetcode.cn id=283 lang=golang
 *
 * [283] 移动零
 */

// @lc code=start
func moveZeroes(nums []int) {
	n := len(nums)
	count := 0
	for i := 0; i < n; i++ {
		if nums[i] != 0 {
			count++
			nums[count-1] = nums[i]
		}
	}
	for i := count; i < n; i++ {
		nums[i] = 0
	}
}

// @lc code=end
