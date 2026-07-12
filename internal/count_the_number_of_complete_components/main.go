package count_the_number_of_complete_components

import (
	"cmp"
	"fmt"
	"slices"
)

func countCompleteComponents(n int, edges [][]int) int {
	complete := 0
	subGraph := [][]int{}

	for i, edge := range edges {
		if edge[0] > edge[1] {
			edges[i] = []int{edge[1], edge[0]}
		}
	}

	slices.SortStableFunc(edges, func(edgeA []int, edgeB []int) int {
		return cmp.Compare(edgeA[0], edgeB[0]) 
	})

	fmt.Printf("edges: %v\n", edges)

	i := 0
	for i < len(edges) {
		currEdge := edges[i]

		if !inGraph(subGraph, currEdge) {
			if isComplete(subGraph) {
				complete += 1
			}

			subGraph = [][]int{currEdge}
		} else {
			subGraph = append(subGraph, currEdge)
		}

		i++
	}

	if !isComplete(subGraph) {
		complete -= 1
	}

	// by definition, if a node is all by itself (no edges reference it), it is complete
	hasEdge := make(map[int]struct{})
	for _, edge := range edges {
		hasEdge[edge[0]] = struct{}{}
		hasEdge[edge[1]] = struct{}{}
	}

	for i := range n {
		_, ok := hasEdge[i]
		if !ok {
			complete += 1
		}
	}

	return complete
}

func isComplete(subGraph [][]int) bool {
	if len(subGraph) <= 1 {
		return true
	}

	edgeMap := make(map[int][]int)
	for _, edge := range subGraph {
		fromNode, toNode := edge[0], edge[1]
		edgeMap[fromNode] = append(edgeMap[fromNode], toNode)
		edgeMap[toNode] = append(edgeMap[toNode], fromNode)
	}

	for nodeA, edgesA := range edgeMap {
		for nodeB := range edgeMap {
			if nodeA != nodeB && !slices.Contains(edgesA, nodeB) {
				return false
			}
		}
	}

	return true
}

func inGraph(subGraph [][]int, checkEdge []int) bool {
	for _, component := range checkEdge {
		for _, edge := range subGraph {
			if edge[0] == component || edge[1] == component {
				return true
			}
		}
	}

	return false
}
