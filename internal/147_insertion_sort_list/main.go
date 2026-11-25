// Package insertsortlist sorts a linked list with insertion sort.
//
// https://leetcode.com/problems/insertion-sort-list/description/
package insertsortlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func insertionSortList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	// elements before curr are "sorted"
	prevI := head
	currI := head.Next
	for currI != nil {
		currINext := currI.Next

		// where do we put currI?
		movedCurrI := false
		if currI.Val < head.Val {
			// if it's before the head, currI is our new head
			currI.Next = head
			head = currI
			movedCurrI = true
		} else {
			// find a spot for currI otherwise
			prevJ := head
			currJ := head.Next
			for currJ != nil && currJ != currI {
				if currI.Val < currJ.Val {
					// insert currI between prevJ and currJ
					prevJ.Next = currI
					currI.Next = currJ
					movedCurrI = true
					break
				}
				prevJ = currJ
				currJ = currJ.Next
			}
		}

		if movedCurrI {
			// make sure we remove currI
			prevI.Next = currINext
		} else {
			prevI = currI
		}

		currI = currINext
	}

	return head
}
