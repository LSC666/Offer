package code

func isMatch(s string, p string) bool {
	lenS := len(s)
	lenP := len(p)
	f := make([][]bool, lenS+1)
	for i := 0; i < lenS+1; i++ {
		f[i] = make([]bool, lenP+1)
	}
	//f创建之后元素的默认值为false
	f[0][0] = true
	for i := 0; i <= lenS; i++ {
		for j := 0; j <= lenP; j++ {
			if j == 0 { //空正则
				f[i][j] = i == 0
			} else { //非空正则
				if p[j-1] != '*' { //p的第j位非*
					if i > 0 && (s[i-1] == p[j-1] || p[j-1] == '.') {
						f[i][j] = f[i-1][j-1]
					}
				} else { //p的第j位是*（此时要把j-1位与j位看为一个整体）
					if j >= 2 { //不看这个整体
						f[i][j] = f[i][j-2]
					} //看这个整体
					if i >= 1 && j >= 2 && (s[i-1] == p[j-2] || p[j-2] == '.') {
						f[i][j] = f[i-1][j]
					}
				}
			}
		}
	}
	return f[lenS][lenP]
}
