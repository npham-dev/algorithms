package count_the_number_of_complete_components

import (
	"fmt"
	"slices"
)

func countCompleteComponents(n int, edges [][]int) int {
	edgeMap := make(map[int][]int)
	for _, edge := range edges {
		fromNode, toNode := edge[0], edge[1]
		edgeMap[fromNode] = append(edgeMap[fromNode], toNode)
		edgeMap[toNode] = append(edgeMap[toNode], fromNode)
	}

	visited := make(map[int]struct{})
	complete := 0

	subGraphs := [][]int{}

	// no edges associated with node, so it's alone and by definition complete
	for i := range n {
		_, ok := visited[i]
		if ok {
			// we already explored this node, skip it
			continue
		}

		_, ok = edgeMap[i]
		if !ok {
			// node has no edges associatd with it, by definition is complete
			visited[i] = struct{}{}
			complete += 1
			continue
		}

		// start exploring subgraphs starting from node i
		subGraphs = append(subGraphs, createSubgraph(&edgeMap, &visited, i))
	}

	fmt.Printf("%v\n", subGraphs)

	for _, graph := range subGraphs {
		if isComplete(&edgeMap, graph) {
			fmt.Printf("  complete: %v\n", graph)
			complete += 1
		}
	}

	return complete
}

func createSubgraph(edgeMap *map[int][]int, visited *map[int]struct{}, start int) []int {
	subgraph := []int{}
	stack := []int{start}
	node := start
	for len(stack) > 0 {
		node, stack = stack[len(stack)-1], stack[:len(stack)-1]
		if _, ok := (*visited)[node]; ok {
			// already visited node
			continue
		}

		// visit node by appending to subgraph
		subgraph = append(subgraph, node)
		(*visited)[node] = struct{}{}

		// add more to stack
		for _, newNode := range (*edgeMap)[node] {
			if _, ok := (*visited)[newNode]; !ok {
				stack = append(stack, newNode)
			}
		}
	}

	return subgraph
}

func isComplete(edgeMap *map[int][]int, subgraph []int) bool {
	for i, node := range subgraph {
		for j := i + 1; j < len(subgraph); j++ {
			if !slices.Contains((*edgeMap)[node], subgraph[j]) {
				return false
			}
		}
	}
	return true
}
