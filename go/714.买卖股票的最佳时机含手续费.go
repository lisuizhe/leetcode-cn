package leetcode

/*
 * @lc app=leetcode.cn id=714 lang=golang
 *
 * [714] 买卖股票的最佳时机含手续费
 */

// @lc code=start
// dp:
// dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
// dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i]-fee
// base:
// dp[-1][0] = 0
// dp[-1][1] = -Inf
func maxProfit(prices []int, fee int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	dp0, dp1 := 0, -prices[0]-fee
	for i := 1; i < n; i++ {
		dp0 = max(dp0, dp1+prices[i])
		dp1 = max(dp1, dp0-prices[i]-fee)
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
