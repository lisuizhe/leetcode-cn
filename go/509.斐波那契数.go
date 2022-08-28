package leetcode

/*
 * @lc app=leetcode.cn id=509 lang=golang
 *
 * [509] 斐波那契数
 */

// @lc code=start
func fib(n int) int {
	if n <= 1 {
		return n
	} else if n == 2 {
		return 1
	}
	current, prev1, prev2 := -1, 1, 1
	for i := 3; i <= n; i++ {
		current = prev1 + prev2
		prev2 = prev1
		prev1 = current
	}
	return current
}

// @lc code=end
