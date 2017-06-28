package main

import (
	"fmt"
)

type node uint

type edge struct {
	first  node
	second node
}

// The quick-find algorithm executes at least M*N instructions, where N is the
// number of nodes in the problem, and M is the number of union operations.
func QuickFind(N int, edges []edge) {
	id := make([]int, N, N)
	for i := 0; i < N; i++ {
		id[i] = i
	}
	for _, e := range edges {
		t := id[e.first]
		if t == id[e.second] {
			// These nodes have already been connected.
			continue
		} else {
			for i := range id {
				if id[i] == t {
					id[i] = id[e.second]
				}
			}
			fmt.Printf("%d-%d\n", e.first, e.second)
		}
	}
	fmt.Printf("\n")
}

func main() {
	input := []edge{
		edge{0, 2},
		edge{1, 4},
		edge{2, 5},
		edge{3, 6},
		edge{0, 4},
		edge{6, 0},
		edge{1, 3},
	}
	QuickFind(len(input), input)

	input2 := []edge{
		edge{3, 4},
		edge{4, 9},
		edge{8, 0},
		edge{2, 3},
		edge{5, 6},
		edge{2, 9},
		edge{5, 9},
		edge{7, 3},
		edge{4, 8},
		edge{5, 6},
		edge{6, 1},
	}
	QuickFind(len(input2), input2)
}
