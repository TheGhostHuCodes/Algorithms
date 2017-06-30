package main

// WeightedQuickUnionWithPathCompressionByHalving executes unions quickly by
// updating a single connection to form the union of two disconnected sets of
// nodes and ensures that the smaller set is joined to the larger tree to
// minimize the time it takes to find the root of a tree.
func WeightedQuickUnionWithPathCompressionByHalving(N int, edges []edge) []edge {
	id := make([]node, N)
	tree_size := make([]int, N)
	for i := 0; i < N; i++ {
		id[i] = node(i)
		tree_size[i] = 1
	}
	var result []edge
	for _, e := range edges {
		i := findRootWithHalving(e.First, id)
		j := findRootWithHalving(e.Second, id)
		// These nodes are not yet connected.
		if i != j {
			// Union the smaller tree to the larger tree.
			if tree_size[i] < tree_size[j] {
				id[i] = j
				tree_size[j] += tree_size[i]
			} else {
				id[j] = i
				tree_size[i] += tree_size[j]
			}
			result = append(result, e)
		}
	}
	return result
}

// findRootWithHalving finds the root of a given node in a set of trees and
// returns the root node. A root node is a node that points to itself. It has
// the side effect of halving the paths of nodes by connecting a node to its
// parent's parent node.
func findRootWithHalving(n node, trees []node) node {
	for {
		if n == trees[n] {
			return n
		} else {
			n = trees[n]
			trees[n] = trees[trees[n]]
		}
	}
}
