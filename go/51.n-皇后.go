package leetcode

/*
 * @lc app=leetcode.cn id=51 lang=golang
 *
 * [51] N 皇后
 */

// @lc code=start
func solveNQueens(n int) [][]string {
	if n <= 0 {
		return [][]string{}
	}
	usedCols := make([]bool, n)
	usedDia1 := make([]bool, 2*n-1)
	usedDia2 := make([]bool, 2*n-1)
	rows := []int{}
	res := [][]string{}
	backtrack(n, &usedCols, &usedDia1, &usedDia2, rows, &res)
	return res
}

func backtrack(n int, usedCols, usedDia1, usedDia2 *[]bool, rows []int, res *[][]string) {
	row := len(rows)
	if row == n {
		*res = append(*res, generateBoard(n, rows))
		return
	}

	for col := 0; col < n; col++ {
		if (*usedCols)[col] || (*usedDia1)[row+col] || (*usedDia2)[row-col+n-1] {
			continue
		}
		(*usedCols)[col] = true
		(*usedDia1)[row+col] = true
		(*usedDia2)[row-col+n-1] = true
		rows = append(rows, col)
		backtrack(n, usedCols, usedDia1, usedDia2, rows, res)
		rows = rows[:row]
		(*usedDia2)[row-col+n-1] = false
		(*usedDia1)[row+col] = false
		(*usedCols)[col] = false
	}
}

func generateBoard(n int, rows []int) []string {
	board := []string{}
	for i := 0; i < n; i++ {
		row := make([]rune, n)
		for j := 0; j < n; j++ {
			if j == rows[i] {
				row[j] = 'Q'
			} else {
				row[j] = '.'
			}
		}
		board = append(board, string(row))
	}
	return board
}

// @lc code=end
