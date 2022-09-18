package leetcode

import "sort"

/*
 * @lc app=leetcode.cn id=698 lang=golang
 *
 * [698] 划分为k个相等的子集
 */

// @lc code=start
func canPartitionKSubsets(nums []int, k int) bool {
	n := len(nums)
	if n < k {
		return false
	}
	sum := 0
	for i := 0; i < n; i++ {
		sum += nums[i]
	}
	if sum%k != 0 {
		return false
	}
	target := sum / k
	// sort from bigger to smaller will increase speed of backtrack
	sort.Slice(nums, func(i, j int) bool { return nums[i] > nums[j] })
	if nums[0] > target {
		return false
	}

	subsets := make([]int, k)
	return backtrack(&nums, &subsets, 0, target, n, k)
}

func backtrack(nums, subsets *[]int, index, target, n, k int) bool {
	if index == n {
		for i := 0; i < k; i++ {
			if (*subsets)[i] != target {
				return false
			}
		}
		return true
	}

	for i := 0; i < k; i++ {
		if (*subsets)[i]+(*nums)[index] > target {
			continue
		}
		// important: already computed
		if i > 0 && (*subsets)[i] == (*subsets)[i-1] {
			continue
		}

		(*subsets)[i] += (*nums)[index]
		if backtrack(nums, subsets, index+1, target, n, k) {
			return true
		}
		(*subsets)[i] -= (*nums)[index]

	}

	return false
}

// @lc code=end
