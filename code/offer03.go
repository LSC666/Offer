package code

//时间复杂度为O(n)，空间复杂度为O(1)
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
