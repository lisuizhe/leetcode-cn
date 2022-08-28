package leetcode

/*
 * @lc app=leetcode.cn id=46 lang=golang
 *
 * [46] 全排列
 */

// @lc code=start
func permute(nums []int) [][]int {
	n := len(nums)
	if n == 0 {
		return [][]int{}
	}
	res := [][]int{}
	used := make([]bool, n)
	track := []int{}
	backtrack(n, nums, &used, track, &res)
	return res
}

func backtrack(n int, nums []int, used *[]bool, track []int, res *[][]int) {
	if len(track) == n {
		tmp := make([]int, n)
		copy(tmp, track)
		*res = append(*res, tmp)
		return
	}

	for i := 0; i < n; i++ {
		if (*used)[i] {
			continue
		}
		(*used)[i] = true
		track = append(track, nums[i])
		backtrack(n, nums, used, track, res)
		track = (track)[:len(track)-1]
		(*used)[i] = false
	}
}

// @lc code=end
