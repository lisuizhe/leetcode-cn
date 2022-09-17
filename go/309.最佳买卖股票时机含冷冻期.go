package leetcode

import "math"

/*
 * @lc app=leetcode.cn id=309 lang=golang
 *
 * [309] 最佳买卖股票时机含冷冻期
 */

// @lc code=start
// dp:
// dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
// dp[i][1] = max(dp[i-1][1], dp[i-2][0]-prices[i])
// base:
// dp[-1][0] = 0
// dp[-1][1] = -Inf
func maxProfit(prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	dp0, dp1, dp0_pre := 0, math.MinInt, 0
	for i := 0; i < n; i++ {
		temp := dp0
		dp0 = max(dp0, dp1+prices[i])
		dp1 = max(dp1, dp0_pre-prices[i])
		dp0_pre = temp
	}
	return dp0
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

// @lc code=end
