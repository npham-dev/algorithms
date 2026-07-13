package concatenation_of_array

import (
	"slices"
	"testing"
)

func TestExamples(t *testing.T) {
	type Example struct {
		Input  []int
		Output []int
	}

	examples := []Example{
		{
			Input:  []int{1, 4, 1, 2},
			Output: []int{1, 4, 1, 2, 1, 4, 1, 2},
		},
		{
			Input:  []int{22, 21, 20, 1},
			Output: []int{22, 21, 20, 1, 22, 21, 20, 1},
		},
	}

	for i, example := range examples {
		actualOutput := getConcatenation(example.Input)
		if !slices.Equal(example.Output, actualOutput) {
			t.Errorf("example %3d | expected %v, got %v", i+1, example.Output, actualOutput)
		}
	}
}
