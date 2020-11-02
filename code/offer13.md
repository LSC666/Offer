### DFS

```
func movingCount(m int, n int, k int) int {
	visited := make([][]int, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]int, n)
	}
	return dfs(visited, m, n, 0, 0, k)
}
func dfs(visited [][]int, m, n, x, y, k int) int {
	if x >= m || y >= n || Count(x, y) > k || visited[x][y] == 1 {
		return 0
	}
	visited[x][y] = 1
	return 1 + dfs(visited, m, n, x+1, y, k) + dfs(visited, m, n, x, y+1, k)
}
func Count(x, y int) int {
	countX := 0
	for x != 0 {
		countX += x % 10
		x = x / 10
	}
	countY := 0
	for y != 0 {
		countY += y % 10
		y = y / 10
	}
	return countX + countY
}
```
```
go
```