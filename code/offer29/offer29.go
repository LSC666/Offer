package offer29

func spiralOrder(matrix [][]int) []int {
	directions := make([][]int, 4)
	directions[0] = []int{0, 1}
	directions[1] = []int{1, 0}
	directions[2] = []int{0, -1}
	directions[3] = []int{-1, 0}
	m := len(matrix)
	n := len(matrix[0])
	res := make([]int, m*n)
	i := 0
	k := 0
	x, y := 0, 0
	for i < m*n {
		res[i] = matrix[x][y]
		matrix[x][y] = -1
		n_x := x + directions[k][0]
		n_y := y + directions[k][1]
		if n_x >= m || n_y >= n || n_x < 0 || n_y < 0 || matrix[n_x][n_y] == -1 {
			k = (k + 1) % 4
			x += directions[k][0]
			y += directions[k][1]
		} else {
			x = n_x
			y = n_y
		}
		i++
	}
	return res
}
