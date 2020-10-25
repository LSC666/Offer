package code

func replaceSpace(s string) string {
	var res string
	for i:=0;i<len(s);i++ {
		if s[i]==' '{
			res+="%20"
		}else {
			res+=string(s[i])
		}
	}
	return res
}
