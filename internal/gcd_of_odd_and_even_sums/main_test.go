package gcd_of_odd_and_even_sums

import "testing"

func TestExamples(t *testing.T) {
	type Example struct {
		Input  int
		Output int
	}

	examples := []Example{
		{
			Input:  4,
			Output: 4,
		},
		{
			Input:  5,
			Output: 5,
		},
	}

	for i, example := range examples {
		actualOutput := gcdOfOddEvenSums(example.Input)
		if actualOutput != example.Output {
			t.Errorf("example %d: expected %d, got %d", i, example.Output, actualOutput)
		}
	}
}
