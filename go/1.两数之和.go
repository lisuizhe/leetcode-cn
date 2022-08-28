package leetcode

/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	n := len(nums)
	if n < 2 {
		return nil
	}

	used := make(map[int]int)
	for i := 0; i < n; i++ {
		if j, ok := used[target-nums[i]]; ok {
			return []int{j, i}
		}
		used[nums[i]] = i
	}
	return nil
}

// @lc code=end
