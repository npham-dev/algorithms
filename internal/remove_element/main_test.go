package remove_element

import (
	"slices"
	"testing"
)

func TestExamples(t *testing.T) {
	type Example struct {
		InputNums []int
		InputVal  int
		Output    []int // length of expected output is equal to k
	}

	examples := []Example{
		{
			InputNums: []int{3, 2, 2, 3},
			InputVal:  3,
			Output:    []int{2, 2},
		},
		{
			InputNums: []int{0, 1, 2, 2, 3, 0, 4, 2},
			InputVal:  2,
			Output:    []int{0, 1, 3, 0, 4},
		},
	}

	for i, example := range examples {
		inputNums := slices.Clone(example.InputNums)
		output := removeElement(inputNums, example.InputVal)

		if output != len(example.Output) {
			t.Errorf("example %3d | expected %d, got %d", i+1, len(example.Output), output)
		}

		inputNums = inputNums[:output]
		slices.Sort(inputNums)

		slices.Sort(example.Output)

		if !slices.Equal(inputNums, example.Output) {
			t.Errorf("example %3d | expected %v, got %v", i+1, example.Output, inputNums)
		}
	}
}
