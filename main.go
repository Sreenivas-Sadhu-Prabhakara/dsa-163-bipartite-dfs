package main

import "fmt"

// IsBipartite reports whether the undirected graph is 2-colorable, by DFS:
// each vertex pushes the opposite color onto unvisited neighbors and
// rejects same-colored ones.
func IsBipartite(n int, edges [][2]int) bool {
	adj := make([][]int, n)
	for _, e := range edges {
		adj[e[0]] = append(adj[e[0]], e[1])
		adj[e[1]] = append(adj[e[1]], e[0])
	}
	color := make([]int, n) // 0 uncolored; 1, 2 the colors
	var paint func(v, c int) bool
	paint = func(v, c int) bool {
		color[v] = c
		for _, w := range adj[v] {
			switch color[w] {
			case 0:
				if !paint(w, 3-c) {
					return false
				}
			case c:
				return false
			}
		}
		return true
	}
	for s := 0; s < n; s++ {
		if color[s] == 0 && !paint(s, 1) {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(IsBipartite(4, [][2]int{{0, 1}, {1, 2}, {2, 3}, {3, 0}}))
	fmt.Println(IsBipartite(3, [][2]int{{0, 1}, {1, 2}, {2, 0}}))
}
