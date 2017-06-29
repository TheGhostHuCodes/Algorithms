package main

type node uint

type edge struct {
	First  node
	Second node
}

type graph struct {
	Input  []edge
	Result []edge
}

func edgesEqual(a []edge, b []edge) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
