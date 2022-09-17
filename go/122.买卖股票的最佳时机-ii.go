package leetcode

/*
 * @lc app=leetcode.cn id=122 lang=golang
 *
 * [122] 买卖股票的最佳时机 II
 */

// @lc code=start
// dp:
// dp[i][j][0] = max(dp[i-1][j][0], dp[i-1][j][1]+prices[i])
// dp[i][j][1] = max(dp[i-1][j][1], dp[i-1][j-1][0]-prices[i])
// (j == Inf) =>
// dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
// dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
//
// base:
// dp[-1][j][0] = dp[i][0][0] = 0
// dp[-1][j][1] = dp[i][0][1] = -Inf
func maxProfit(prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}

	dp := make([][2]int, n)
	for i := 0; i < n; i++ {
		var dp0, dp1 int
		if i == 0 {
			dp0 = 0
			dp1 = -prices[i]
		} else {
			dp0 = max(dp[i-1][0], dp[i-1][1]+prices[i])
			dp1 = max(dp[i-1][1], dp[i-1][0]-prices[i])
		}
		dp[i] = [2]int{dp0, dp1}
	}
	return dp[n-1][0]
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

// @lc code=end
