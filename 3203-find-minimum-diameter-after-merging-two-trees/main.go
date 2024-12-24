package main

type Solution struct {
	dia int
}

func minimumDiameterAfterMerge(edges1, edges2 [][]int) int {
	s := Solution{}
	n1, n2 := len(edges1)+1, len(edges2)+1

	graph1, graph2 := make([][]int, n1), make([][]int, n2)

	for _, edge := range edges1 {
		graph1[edge[0]] = append(graph1[edge[0]], edge[1])
		graph1[edge[1]] = append(graph1[edge[1]], edge[0])
	}

	for _, edge := range edges2 {
		graph2[edge[0]] = append(graph2[edge[0]], edge[1])
		graph2[edge[1]] = append(graph2[edge[1]], edge[0])
	}

	s.dia = -1
	visited1 := make([]bool, n1)
	s.getDia(0, graph1, visited1)
	d1 := s.dia

	s.dia = -1
	visited2 := make([]bool, n2)
	s.getDia(0, graph2, visited2)
	d2 := s.dia

	cand := (d1+1)/2 + (d2+1)/2 + 1
	return max(max(cand, d1), d2)
}

func (s *Solution) getDia(src int, graph [][]int, visited []bool) int {
	visited[src] = true
	dch, sdch := -1, -1

	for _, child := range graph[src] {
		if !visited[child] {
			ch := s.getDia(child, graph, visited)
			if ch > dch {
				sdch = dch
				dch = ch
			} else if ch > sdch {
				sdch = ch
			}
		}
	}

	if dch+sdch+2 > s.dia {
		s.dia = dch + sdch + 2
	}
	return dch + 1
}
