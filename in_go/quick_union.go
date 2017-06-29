package main

// QuickUnion executes unions quickly by updating a single connection to form
// the union of two disconnected sets of nodes.
func QuickUnion(N int, edges []edge) []edge {
	id := make([]node, N, N)
	for i := 0; i < N; i++ {
		id[i] = node(i)
	}
	var result []edge
	for _, e := range edges {
		i := findRoot(e.First, id)
		j := findRoot(e.Second, id)
		// These nodes are not yet connected.
		if i != j {
			id[i] = j
			result = append(result, e)
		}
	}
	return result
}

// findRoot finds the root of a given node in a set of trees and returns the
// root node. A root node is a node that points to itself.
func findRoot(n node, trees []node) node {
	for {
		if n == trees[n] {
			return n
		} else {
			n = trees[n]
		}
	}
}
