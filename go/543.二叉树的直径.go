package leetcode

/*
 * @lc app=leetcode.cn id=543 lang=golang
 *
 * [543] 二叉树的直径
 */

// @lc code=start
//
// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)
	diameterRoot := leftDepth + rightDepth
	diameterLeft := diameterOfBinaryTree(root.Left)
	diameterRight := diameterOfBinaryTree(root.Right)
	return max(diameterRoot, max(diameterLeft, diameterRight))
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

// @lc code=end
