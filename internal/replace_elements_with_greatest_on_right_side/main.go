package replace_elements_with_greatest_on_right_side

func replaceElements(arr []int) []int {
	maxRight := -1
	for i := len(arr) - 1; i >= 0; i-- {
		t := arr[i]
		arr[i] = maxRight
		maxRight = max(maxRight, t)
	}

	return arr
}
