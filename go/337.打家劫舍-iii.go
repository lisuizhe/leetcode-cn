package leetcode

import "math"

/*
 * @lc app=leetcode.cn id=337 lang=golang
 *
 * [337] 打家劫舍 III
 */

// @lc code=start
//
// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rob(root *TreeNode) int {
	return max(subrob(root))
}

func subrob(root *TreeNode) (int, int) {
	if root == nil {
		return 0, math.MinInt
	}
	left0, left1 := subrob(root.Left)
	right0, right1 := subrob(root.Right)
	return max(left0, left1) + max(right0, right1), left0 + right0 + root.Val
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

// @lc code=end
