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
	return 1 + dfs(visited, m, n, x+1, y, k) + dfs(visited, m, n, x, y+1, k)//只需考虑向右和向下
}
//计算x、y的数位之和
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
### BFS
```
type pair struct {
	x, y int
}

func movingCount(m int, n int, k int) int {
	visited := make([][]int, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]int, n)
	}
	queue := []pair{}
	queue = append(queue, pair{0, 0})
	res := 1
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		x, y := cur.x, cur.y
		if x >= m || y >= n || Count(x, y) > k || visited[x][y] == 1 {
			continue
		}
		visited[x][y] = 1
		res++
        //只需考虑向右和向下
		queue = append(queue, pair{x + 1, y})
		queue = append(queue, pair{x, y + 1})
	}
	return res
}

//计算x、y的数位之和
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