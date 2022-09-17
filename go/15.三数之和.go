package leetcode

import (
	"fmt"
)

/*
 * @lc app=leetcode.cn id=15 lang=golang
 *
 * [15] 三数之和
 */

// @lc code=start
var empty struct{}

func threeSum(nums []int) [][]int {
	res := [][]int{}
	n := len(nums)
	// sort.Slice(nums, func(a, b int) bool {
	// 	return a < b
	// })
	nums = quickSort(nums)
	used := make(map[string]struct{})
	for i := 0; i < n-2; i++ {
		lo, hi := i+1, n-1
		for lo < hi {
			sum := nums[i] + nums[lo] + nums[hi]
			cur := []int{nums[i], nums[lo], nums[hi]}
			if sum == 0 {
				key := fmt.Sprintf("%d:%d:%d", cur[0], cur[1], cur[2])
				if _, ok := used[key]; !ok {
					res = append(res, cur)
					used[key] = empty
				}
				lo++
				hi--
			} else if sum < 0 {
				lo++
			} else {
				hi--
			}
		}
	}
	return res
}

func quickSort(nums []int) []int {
	n := len(nums)
	if n == 0 {
		return []int{}
	}
	newNums := make([]int, n)
	for i := 0; i < n; i++ {
		newNums[i] = nums[i]
	}
	partitionSort(newNums, 0, n-1)
	return newNums
}

func partitionSort(nums []int, start, end int) {
	if end-start < 1 {
		return
	}
	pivot := nums[end]
	splitIdx := start
	for i := start; i < end; i++ {
		if nums[i] < pivot {
			nums[i], nums[splitIdx] = nums[splitIdx], nums[i]
			splitIdx++
		}
	}
	nums[end] = nums[splitIdx]
	nums[splitIdx] = pivot
	partitionSort(nums, start, splitIdx-1)
	partitionSort(nums, splitIdx+1, end)
}

// @lc code=end
