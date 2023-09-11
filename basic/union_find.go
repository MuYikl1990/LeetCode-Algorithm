package basic

func find(nodes map[int]int, x int) int {
	root := x
	for root != nodes[root] {
		root = nodes[root]
	}

	for x != root {
		tmp := nodes[x]
		nodes[x] = root
		x = tmp
	}
	return root
}

func union(x, y int, nodes map[int]int) {
	xRoot, yRoot := find(nodes, x), find(nodes, y)
	if xRoot != yRoot {
		nodes[xRoot] = yRoot
	}
}

func add(x int, nodes map[int]int) {
	if _, ok := nodes[x]; !ok {
		nodes[x] = x
	}
}
