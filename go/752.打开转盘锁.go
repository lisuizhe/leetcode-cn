package leetcode

/*
 * @lc app=leetcode.cn id=752 lang=golang
 *
 * [752] 打开转盘锁
 */

// @lc code=start
func openLock(deadends []string, target string) int {
	var empty struct{}
	if len(target) != 4 {
		return -1
	}
	for _, r := range target {
		if r > '9' || r < '0' {
			return -1
		}
	}

	deads := make(map[string]struct{})
	for _, d := range deadends {
		deads[d] = empty
	}
	q1 := map[string](struct{}){}
	q2 := map[string](struct{}){}
	visited := map[string](struct{}){}
	q1["0000"] = empty
	q2[target] = empty
	step := 0

	for len(q1) != 0 && len(q2) != 0 {
		// optimization to BFS from smaller set
		if len(q2) < len(q1) {
			tmp := q1
			q1 = q2
			q2 = tmp
		}

		tmp := map[string](struct{}){}
		for current := range q1 {
			if _, dead := deads[current]; dead {
				continue
			}
			if _, currentVisited := q2[current]; currentVisited {
				return step
			}
			visited[current] = empty
			for j := 0; j < 4; j++ {
				up := plusOne(current, j)
				if _, upVisited := visited[up]; !upVisited {
					tmp[up] = empty
				}
				down := minusOne(current, j)
				if _, downVistited := visited[down]; !downVistited {
					tmp[down] = empty
				}
			}

		}
		step++
		q1 = q2
		q2 = tmp
	}

	return -1
}

func plusOne(s string, i int) string {
	r := []rune(s)
	if r[i] == '9' {
		r[i] = '0'
	} else {
		r[i] = r[i] + 1
	}
	return string(r)
}

func minusOne(s string, i int) string {
	r := []rune(s)
	if r[i] == '0' {
		r[i] = '9'
	} else {
		r[i] = r[i] - 1
	}
	return string(r)
}

// @lc code=end
