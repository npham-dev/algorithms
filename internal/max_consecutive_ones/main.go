package maxconsecutiveones

func findMaxConsecutiveOnes(nums []int) int {
	maxCount := 0
	count := 0

	for _, n := range nums {
		if n == 1 {
			count += 1
		} else {
			maxCount = max(maxCount, count)
			count = 0
		}
	}

	return max(maxCount, count)
}
