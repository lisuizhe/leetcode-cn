package leedcode

/*
 * @lc app=leetcode.cn id=130 lang=golang
 *
 * [130] 被围绕的区域
 */

// @lc code=start
var dirs = [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

func solve(board [][]byte) {
	for i := range board {
		for j := range board[i] {
			if (i == 0 || i == len(board)-1 || j == 0 || j == len(board[i])-1) && board[i][j] == 'O' {
				dfs(i, j, board)
			}
		}
	}
	for i := range board {
		for j := range board[i] {
			if board[i][j] == '#' {
				board[i][j] = 'O'
			} else if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}
}

func dfs(i, j int, board [][]byte) {
	if i < 0 || i >= len(board) || j < 0 || j >= len(board[i]) {
		return
	}
	if board[i][j] == 'O' {
		board[i][j] = '#'
		for _, dir := range dirs {
			dfs(i+dir[0], j+dir[1], board)
		}
	}
}

// @lc code=end
