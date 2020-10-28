package code

/*
递归
 */
func fib(n int) int {
	if n==0{
		return 0
	}
	if n==1{
		return 1
	}
	return fib(n-1)+fib(n-2)
}

/*
循环
 */
func func10_2(n int) int {
	if n==0{
		return 0
	}
	if n==1{
		return 1
	}
	t,a,b:=0,0,1
	for i:=2;i<=n;i++ {
		t=(a+b)%1000000007
		a=b
		b=t
	}
	return t
}