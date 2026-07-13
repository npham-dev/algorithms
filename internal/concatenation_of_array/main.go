package concatenation_of_array

func getConcatenation(nums []int) []int {
	l := len(nums)
	result := make([]int, l*2)
	for i := range l {
		result[i] = nums[i]
		result[l+i] = nums[i]
	}
	return result
}
