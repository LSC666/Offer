package code

/*
回溯法(也可以说是DFS+剪枝)
这种方法会破坏原数组
*/
func exist(board [][]byte, word string) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if dfs(board, word, i, j, 0) {
				return true
			}
		}
	}
	return false
}
func dfs(board [][]byte, word string, i, j, index int) bool {
	if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) || board[i][j] != word[index] {
		return false
	}
	if index == len(word)-1 {
		return true
	}
	tmp := board[i][j]
	board[i][j] = '.'
	res := dfs(board, word, i+1, j, index+1) || dfs(board, word, i-1, j, index+1) || dfs(board, word, i, j+1, index+1) || dfs(board, word, i, j-1, index+1)
	board[i][j] = tmp
	return res
}
