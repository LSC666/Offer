package offer32

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	queue := []*TreeNode{}
	queue = append(queue, root)
	res := [][]int{}
	temp := []int{}
	pre := root
	direction := 0 //0为左，1为右
	for len(queue) != 0 {
		t := queue[0]
		queue = queue[1:]
		temp = append(temp, t.Val)
		if t.Left != nil {
			queue = append(queue, t.Left)
		}
		if t.Right != nil {
			queue = append(queue, t.Right)
		}
		if pre == t {
			if direction == 0 {
				res = append(res, temp)
				direction = 1
			} else {
				tt := []int{}
				for i := len(temp) - 1; i >= 0; i-- {
					tt = append(tt, temp[i])
				}
				res = append(res, tt)
				direction = 0
			}
			temp = []int{}
			if len(queue) != 0 {
				pre = queue[len(queue)-1]
			}
		}
	}
	return res
}
