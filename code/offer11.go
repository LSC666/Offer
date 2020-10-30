package code

/*
从头到尾遍历
*/
func minArray(numbers []int) int {
	for i := 0; i < len(numbers)-1; i++ {
		if numbers[i] > numbers[i+1] {
			return numbers[i+1]
		}
	}
	return numbers[0]
}

/*
二分查找
*/
func func11_2(numbers []int) int {
	l, r := 0, len(numbers)-1
	for l < r {
		m := (l + r) / 2
		if numbers[m] < numbers[r] { //中间元素位于最小元素的右侧或者就是最小元素
			r = m
		} else if numbers[m] > numbers[r] { //中间元素位于最小元素的左侧
			l = m + 1
		} else {
			r--
		}
	}
	return numbers[l]
}
