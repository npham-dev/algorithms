package count_the_number_of_complete_components

import "testing"

func TestExamples(t *testing.T) {
	type Example struct {
		InputN     int
		InputEdges [][]int
		Output     int
	}

	examples := []Example{
		{
			InputN:     6,
			InputEdges: [][]int{{0, 1}, {0, 2}, {1, 2}, {3, 4}},
			Output:     3,
		},
		{
			InputN:     6,
			InputEdges: [][]int{{0, 1}, {0, 2}, {1, 2}, {3, 4}, {3, 5}},
			Output:     1,
		},
		{
			InputN:     6,
			InputEdges: [][]int{{0, 1}, {2, 3}, {4, 5}, {0, 2}, {4, 0}, {5, 3}, {0, 3}, {0, 5}, {1, 2}, {1, 3}, {1, 4}, {1, 5}, {2, 4}, {2, 5}, {3, 4}},
			Output:     1,
		},
		{
			InputN:     4,
			InputEdges: [][]int{{1, 0}, {3, 0}, {3, 2}},
			Output:     0,
		},
		{
			InputN:     4,
			InputEdges: [][]int{{2, 0}, {3, 1}, {3, 2}},
			Output:     0,
		},
	}
	for i, example := range examples {
		actualOutput := countCompleteComponents(example.InputN, example.InputEdges)
		if example.Output != actualOutput {
			t.Errorf("example %3d | expected %d, got %d", i+1, example.Output, actualOutput)
		}
	}
}

func TestInGraph(t *testing.T) {
	type Example struct {
		InputGraph [][]int
		InputEdge  []int
		Output     bool
	}

	examples := []Example{
		{
			InputGraph: [][]int{{3, 4}, {1, 4}},
			InputEdge:  []int{2, 3},
			Output:     true,
		},
		{
			InputGraph: [][]int{{3, 4}, {1, 4}, {2, 3}},
			InputEdge:  []int{2, 3},
			Output:     true,
		},
		{
			InputGraph: [][]int{},
			InputEdge:  []int{2, 3},
			Output:     false,
		},
	}
	for i, example := range examples {
		actualOutput := inGraph(example.InputGraph, example.InputEdge)
		if example.Output != actualOutput {
			t.Errorf("example %3d | expected %v, got %v", i+1, example.Output, actualOutput)
		}
	}
}
