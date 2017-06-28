package main

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
