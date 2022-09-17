package leetcode

import "math"

/*
 * @lc app=leetcode.cn id=123 lang=golang
 *
 * [123] 买卖股票的最佳时机 III
 */

// @lc code=start
// dp:
// dp[i][j][0] = max(dp[i-1][j][0], dp[i-1][j][1]+prices[i])
// dp[i][j][1] = max(dp[i-1][j][1], dp[i-1][j-1][0]-prices[i])
// base:
// dp[-1][j][0] = dp[i][0][0] = 0
// dp[-1][j][1] = dp[i][0][1] = -Inf
func maxProfit(prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}

	dp := make([][3][2]int, n)
	for i := 0; i < n; i++ {
		dp[i] = [3][2]int{}
		for j := 0; j <= 2; j++ {
			dp[i][j] = [2]int{}
			if j == 0 {
				dp[i][j] = [2]int{0, math.MinInt}
			}
		}
	}
	for i := 0; i < n; i++ {
		for j := 1; j <= 2; j++ {
			if i == 0 {
				dp[i][j] = [2]int{0, -prices[i]}
			} else {
				dp[i][j] = [2]int{
					max(dp[i-1][j][0], dp[i-1][j][1]+prices[i]),
					max(dp[i-1][j][1], dp[i-1][j-1][0]-prices[i]),
				}
			}
		}
	}
	return dp[n-1][2][0]
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

// @lc code=end
