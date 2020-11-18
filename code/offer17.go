package code

func printNumbers(n int) []int {
	length:=10
	for i:=0;i<n-1;i++ {
		length*=10
	}
	res:=make([]int,length-1)
	for i:=0;i<length-1;i++ {
		res[i]=i+1
	}
	return res
}
func overflow(s []int) bool {
	isOverFlow:=false
	var carry int = 0
	for i:=len(s)-1;i>=0;i-- {
		cur:=s[i]+carry
		if i==len(s)-1{
			cur++
		}
		if cur>=10{
			if i==0{
				isOverFlow=true
			}else {
				carry=1
				s[i]=cur-10
			}
		}else {
			s[i]=cur
			break
		}
	}
	return isOverFlow
}
