package main

type Node133 struct {
    Val int
    Neighbors []*Node133
}

func cloneGraph(node *Node133) *Node133 {
	visited := make(map[*Node133]*Node133)
	return dfs133(node, visited)
}

func dfs133(node *Node133, visited map[*Node133]*Node133) *Node133 {
	if node == nil {
		return node
	}
	if _, ok := visited[node]; ok {
		return visited[node]
	}
	cur := &Node133{Val: node.Val}
	visited[node] = cur
	for _, neighbor := range node.Neighbors {
		cur.Neighbors = append(cur.Neighbors, dfs133(neighbor, visited))
	}
	return cur
}
