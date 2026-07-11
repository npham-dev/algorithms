package maxconsecutiveones

import "testing"

func TestExamples(t *testing.T) {
	type Example struct {
		Input  []int
		Output int
	}

	examples := []Example{
		{
			Input:  []int{1, 1, 0, 1, 1, 1},
			Output: 3,
		},
		{
			Input:  []int{1, 0, 1, 1, 0, 1},
			Output: 2,
		},
	}

	for i, example := range examples {
		output := findMaxConsecutiveOnes(example.Input)
		if example.Output != output {
			t.Errorf("example %03d | expected %d, got %d", i+1, example.Output, output)
		}
	}
}
