package offer27

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	t := root.Right
	root.Right = root.Left
	root.Left = t
	mirrorTree(root.Left)
	mirrorTree(root.Right)
	return root
}
