package ofer31

func validateStackSequences(pushed []int, popped []int) bool {
	stack := []int{}
	flag := make(map[int]bool) //标识某元素是否已入栈
	for _, v := range pushed {
		flag[v] = false
	}
	for _, pop := range popped {
		if len(stack) == 0 || stack[len(stack)-1] != pop {
			for _, push := range pushed {
				if push != pop {
					if flag[push] == false {
						stack = append(stack, push)
						flag[push] = true
					}
				} else {
					if flag[pop] == true {
						return false
					}
					flag[pop] = true
					break
				}
			}
		} else {
			stack = stack[:len(stack)-1]
		}
	}
	return true
}
