package remove_element

func removeElement(nums []int, val int) int {
	i := 0
	end := len(nums) - 1
	for i <= end {
		if nums[i] == val {
			nums[i] = nums[end]
			end -= 1
		} else {
			i += 1
		}
	}
	return i
}
