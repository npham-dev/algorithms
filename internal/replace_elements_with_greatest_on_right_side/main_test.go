package replace_elements_with_greatest_on_right_side

import (
	"slices"
	"testing"
)

func TestExamples(t *testing.T) {
	type Examples struct {
		Input  []int
		Output []int
	}

	examples := []Examples{
		{
			Input:  []int{2, 4, 5, 3, 1, 2},
			Output: []int{5, 5, 3, 2, 2, -1},
		},
		{
			Input:  []int{3, 3},
			Output: []int{3, -1},
		},
	}

	for i, example := range examples {
		actualOutput := replaceElements(example.Input)
		if !slices.Equal(example.Output, actualOutput) {
			t.Errorf("example %3d | expected %v, got %v", i+1, example.Output, actualOutput)
		}
	}
}
