package code

/*
从右上角或左下角出发都行，这样每次都能排除一行或一列,
左上角和右下角都不行
此处选取右上角出发
*/
func findNumberIn2DArray(matrix [][]int, target int) bool {
	rows := len(matrix)
	if rows == 0 {
		return false
	}
	cols := len(matrix[0])
	row, col := 0, cols-1
	for row < rows && col >= 0 {
		if matrix[row][col] > target {
			col -= 1
		} else if matrix[row][col] < target {
			row += 1
		} else {
			return true
		}
	}
	return false
}
