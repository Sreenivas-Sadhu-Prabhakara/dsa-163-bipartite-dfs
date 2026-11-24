package main

import "testing"

func TestIsBipartite(t *testing.T) {
	cases := []struct {
		n     int
		edges [][2]int
		want  bool
	}{
		{4, [][2]int{{0, 1}, {1, 2}, {2, 3}, {3, 0}}, true},
		{3, [][2]int{{0, 1}, {1, 2}, {2, 0}}, false},
		{7, [][2]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}, {4, 5}, {5, 6}, {6, 0}}, false}, // 7-cycle
		{6, [][2]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}, {4, 5}, {5, 0}}, true},          // 6-cycle
		{4, [][2]int{{0, 1}, {2, 3}}, true},
		{1, nil, true},
	}
	for _, c := range cases {
		if got := IsBipartite(c.n, c.edges); got != c.want {
			t.Errorf("IsBipartite(n=%d, %v) = %v, want %v", c.n, c.edges, got, c.want)
		}
	}
}
