package leetcode

/*
 * @lc app=leetcode.cn id=198 lang=golang
 *
 * [198] 打家劫舍
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
	if n == 0 {
		return 0
	}
	dp0, dp1 := 0, nums[0]
	for i := 1; i < n; i++ {
		dp0_pre := dp0
		dp0 = max(dp0, dp1)
		dp1 = dp0_pre + nums[i]
	}
	return max(dp0, dp1)
}

// func rob1(nums []int) int {
// 	n := len(nums)
// 	if n == 0 {
// 		return 0
// 	}
// 	dp := make([][2]int, n)
// 	for i := 0; i < n; i++ {
// 		if i == 0 {
// 			dp[i] = [2]int{0, nums[i]}
// 		} else {
// 			dp[i] = [2]int{
// 				max(dp[i-1][0], dp[i-1][1]),
// 				dp[i-1][0] + nums[i],
// 			}
// 		}
// 	}
// 	return max(dp[n-1][0], dp[n-1][1])
// }

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

// @lc code=end
