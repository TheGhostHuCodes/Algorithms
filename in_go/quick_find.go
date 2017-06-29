package main

import (
	"fmt"
)

// The quick-find algorithm executes at least M*N instructions, where N is the
// number of nodes in the problem, and M is the number of union operations.
func QuickFind(N int, edges []edge) {
	id := make([]int, N, N)
	for i := 0; i < N; i++ {
		id[i] = i
	}
	for _, e := range edges {
		t := id[e.First]
		if t == id[e.Second] {
			// These nodes have already been connected.
			continue
		} else {
			for i := range id {
				if id[i] == t {
					id[i] = id[e.Second]
				}
			}
			fmt.Printf("%d-%d\n", e.First, e.Second)
		}
	}
	fmt.Printf("\n")
}
