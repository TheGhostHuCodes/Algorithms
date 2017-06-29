package main

// WeightedQuickUnion executes unions quickly by updating a single connection
// to form the union of two disconnected sets of nodes and ensures that the
// smaller set is joined to the larger tree to minimize the time it takes to
// find the root of a tree.
func WeightedQuickUnion(N int, edges []edge) []edge {
	id := make([]node, N)
	tree_size := make([]int, N)
	for i := 0; i < N; i++ {
		id[i] = node(i)
		tree_size[i] = 1
	}
	var result []edge
	for _, e := range edges {
		i := findRoot(e.First, id)
		j := findRoot(e.Second, id)
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
