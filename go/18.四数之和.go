package leetcode

import (
	"fmt"
)

/*
 * @lc app=leetcode.cn id=18 lang=golang
 *
 * [18] 四数之和
 */

// @lc code=start
var empty struct{}

func fourSum(nums []int, target int) [][]int {
	n := len(nums)
	res := [][]int{}
	if n < 4 {
		return res
	}
	quickSort(nums)
	used := map[string]struct{}{}
	for i := 0; i < n-2; i++ {
		for j := i + 1; j < n-1; j++ {
			lo, hi := j+1, n-1
			for lo < hi {
				sum := nums[i] + nums[j] + nums[lo] + nums[hi]
				if sum == target {
					cur := []int{nums[i], nums[j], nums[lo], nums[hi]}
					key := fmt.Sprintf("%d:%d:%d:%d", cur[0], cur[1], cur[2], cur[3])
					if _, ok := used[key]; !ok {
						res = append(res, cur)
						used[key] = empty
					}
					lo++
					hi--
				} else if sum < target {
					lo++
				} else {
					hi--
				}
			}
		}
	}
	return res
}

func quickSort(nums []int) {
	n := len(nums)
	if n == 0 {
		return
	}
	partitionSort(nums, 0, n-1)
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
