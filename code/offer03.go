package code

/*
时间复杂度O(n)，空间复杂度O(1)
从头到尾扫描数组，设现已扫描至下标为i的元素，值为k，判断i是否等于k，若相等，
则继续扫描；若不相等，则寻找下标为k的元素，设其值为m，若k不等于m，则交换这两个元素，
继续从下标为i的位置继续扫描；否则返回k。
*/
func findRepeatNumber(nums []int) int {
	for i := 0; i < len(nums); {
		if i != nums[i] {
			if nums[i] == nums[nums[i]] {
				return nums[i]
			} else {
				t := nums[i]
				nums[i] = nums[t]
				nums[t] = t
			}
		} else {
			i++
		}
	}
	return 0
}

//时间复杂度为O(n)，空间复杂度为O(n)
/*
利用map，数组中的元素在0至n-1之间，可以从头到尾扫描数组，
每扫描到一个元素，便向map中查询该元素是否存在，若已存在，
则判定有重复元素；若不存在，则将该元素加入至map中，继续扫描。
*/
func func2(nums []int) int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if m[nums[i]] != 0 {
			return nums[i]
		}
		m[nums[i]]++
	}
	return 0
}
