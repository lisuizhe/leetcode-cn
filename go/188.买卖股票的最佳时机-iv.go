package leetcode

import "math"

/*
 * @lc app=leetcode.cn id=188 lang=golang
 *
 * [188] 买卖股票的最佳时机 IV
 */

// @lc code=start
// dp:
// dp[i][j][0] = max(dp[i-1][j][0], dp[i-1][j][1]+prices[i])
// dp[i][j][1] = max(dp[i-1][j][1], dp[i-1][j-1][0]-prices[i])
// base:
// dp[-1][j][0] = dp[i][0][0] = 0
// dp[-1][j][1] = dp[i][0][-1] = -Inf
func maxProfit(k int, prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	if k > n/2 {
		return maxProfitKIsInf(prices)
	}

	dp := make([][][2]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([][2]int, k+1)
		for j := 0; j <= k; j++ {
			dp[i][j] = [2]int{}
			if j == 0 {
				dp[i][j] = [2]int{0, math.MinInt}
			}
		}
	}
	for i := 0; i < n; i++ {
		for j := 1; j <= k; j++ {
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
	return dp[n-1][k][0]
}

func maxProfitKIsInf(prices []int) int {
	n := len(prices)
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
