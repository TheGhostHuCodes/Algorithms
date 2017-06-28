package main

import (
	"fmt"
)

// QuickUnion executes unions quickly by updating a single connection to form
// the union of two disconnected sets of nodes.
func QuickUnion(N int, edges []edge) {
	id := make([]node, N, N)
	for i := 0; i < N; i++ {
		id[i] = node(i)
	}
	for _, e := range edges {
		i := findRoot(e.first, id)
		j := findRoot(e.second, id)
		// These nodes are not yet connected.
		if i != j {
			id[i] = j
			fmt.Printf("%d-%d\n", e.first, e.second)
		}
	}
	fmt.Printf("\n")
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
