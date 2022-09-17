package leetcode

/*
 * @lc app=leetcode.cn id=213 lang=golang
 *
 * [213] 打家劫舍 II
 */

// @lc code=start
// dp:
// dp[i][0] = max(dp[i-1][0], dp[i-1][1])
// dp[i][1] = dp[i-1][0]+nums[i]
// base:
// dp[-1][0] = 0
// dp[-1][1] = -Inf
func rob(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	return max(subrob(nums[:n-1]), subrob(nums[1:]))
}

func subrob(nums []int) int {
	dp0, dp1 := 0, nums[0]
	for i := 1; i < len(nums); i++ {
		dp0_pre := dp0
		dp0 = max(dp0, dp1)
		dp1 = dp0_pre + nums[i]
	}
	return max(dp0, dp1)
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

// @lc code=end
