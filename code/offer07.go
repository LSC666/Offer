package code

//Definition for a binary tree node.
 type TreeNode struct {
     Val int
     Left *TreeNode
     Right *TreeNode
 }

 /*
 递归方法
  */
func buildTree(preorder []int, inorder []int) *TreeNode {
	if preorder==nil || len(preorder)==0{
		return nil
	}
	rootIdx :=0//表示根节点在inorder数组中的位置，也表示左子树的节点数
	for i:=range inorder{
		if inorder[i]==preorder[0]{
			rootIdx =i
			break
		}
	}
	root:=&TreeNode{
		Val:   preorder[0],
		Left:  nil,
		Right: nil,
	}
	root.Left=buildTree(preorder[1: rootIdx+1],inorder[:rootIdx])
	root.Right=buildTree(preorder[rootIdx+1:],inorder[rootIdx+1:])
	return root
}