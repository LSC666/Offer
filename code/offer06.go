package code


//Definition for singly-linked list.
type ListNode struct {
    Val int
    Next *ListNode
}

/*
创建无长度的slice，遍历链表，将节点值存入slice中，最后反转slice即可。
 */
func reversePrint(head *ListNode) []int {
	if head==nil{
		return nil
	}
	var res []int
	p:=head
	for p!=nil {
		res=append(res, p.Val)
		p=p.Next
	}
	l:=len(res)
	for i:=0;i<l/2;i++ {
		t:=res[i]
		res[i]=res[l-1-i]
		res[l-1-i]=t
	}
	return res
}

/*
先遍历链表，统计链表长度，根据长度创建slice，
遍历链表，将节点值依次存入slice的末端。
 */
func func06(head *ListNode) []int {
	if head==nil{
		return nil
	}
	l:=0
	p:=head
	for p!=nil {
		l++
		p=p.Next
	}
	res:=make([]int,l)
	p=head
	i:=l-1
	for p!=nil {
		res[i]=p.Val
		p=p.Next
		i--
	}
	return res
}