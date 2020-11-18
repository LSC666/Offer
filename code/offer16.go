package code


func myPow(x float64, n int) float64 {
	if n<0{
		x=1/x
		n=-n
	}
	var res float64=1
	cur:=x
	for i:=n;i!=0;i=i>>1 {
		if i&1==1{
			res=res*cur
		}
		cur=cur*cur
	}
	return res
}
