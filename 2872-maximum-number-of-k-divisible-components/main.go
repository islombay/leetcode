package main

func maxKDivisibleComponents(n int, edges [][]int, values []int, k int) int {
	g := make(map[int]map[int]bool)
	for _, edge := range edges {
		u, v := edge[0], edge[1]
		if g[u] == nil {
			g[u] = make(map[int]bool)
		}
		if g[v] == nil {
			g[v] = make(map[int]bool)
		}
		g[u][v] = true
		g[v][u] = true
	}
	res := 1
	var dfs func(int) int
	dfs = func(node int) int {
		total := values[node]
		for next_node := range g[node] {
			delete(g[next_node], node)
			next_total := dfs(next_node)
			if next_total%k == 0 {
				res++
			} else {
				total += next_total
			}
		}
		return total
	}
	dfs(0)
	return res
}
