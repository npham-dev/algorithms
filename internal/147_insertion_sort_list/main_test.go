package insertsortlist

import (
	"slices"
	"testing"
)

type Case struct {
	Input    []int
	Expected []int
}

func TestDefaultCases(t *testing.T) {
	cases := []Case{
		{
			Input:    []int{4, 3, 2, 1},
			Expected: []int{1, 2, 3, 4},
		},
		{
			Input:    []int{-1, 5, 3, 4, 0},
			Expected: []int{-1, 0, 3, 4, 5},
		},
	}

	for i, c := range cases {
		input := sliceToLinkedList(c.Input)
		output := insertionSortList(input)
		if !slices.Equal(linkedListToSlice(output), c.Expected) {
			t.Errorf("failed case %d, expected %+v but got %+v", i, c.Expected, linkedListToSlice(output))
		}
	}
}

// slice -> linked list
func sliceToLinkedList(list []int) *ListNode {
	if len(list) == 0 {
		return nil
	}

	// empty head
	head := &ListNode{}
	curr := head
	for _, element := range list {
		curr.Next = &ListNode{Val: element}
		curr = curr.Next
	}

	return head.Next
}

func linkedListToSlice(head *ListNode) []int {
	result := []int{}
	curr := head
	for curr != nil {
		result = append(result, curr.Val)
		curr = curr.Next
	}
	return result
}
