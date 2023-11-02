package weighted

func (g *Graph) AddDirectedEdge(from, to, weight int) {
	edge := Edge{
		To:     to,
		Weight: weight,
	}
	g.adjacencyList[from] = append(g.adjacencyList[from], edge)
	//since its an undirected graph, we add the reverse
	// edge as well
	reverseEdge := Edge{
		To:     from,
		Weight: weight,
	}
	g.adjacencyList[to] = append(g.adjacencyList[to], reverseEdge)
}

func (g *Graph) RemoveDirectedEdge(from, to int) {
	//find and remove the edge 'from' to 'to'
	for i, edge := range g.adjacencyList[from] {
		if edge.To == to {
			g.adjacencyList[from] = append(g.adjacencyList[from][:i], g.adjacencyList[from][i+1:]...)
			break
		}
	}
	//remove the reverse edge from 'to' to 'form'
	for i, edge := range g.adjacencyList[to] {
		if edge.To == from {
			g.adjacencyList[to] = append(g.adjacencyList[to][:i], g.adjacencyList[to][1+i:]...)
			break
		}
	}
}

func (g *Graph) RemoveUndirectedVertex(vertex int) {
	//Remove vertex & its edges from the adjacency list
	delete(g.adjacencyList, vertex)
	//Remove all edges pointing to the removed vertex
	for i, edge := range g.adjacencyList {
		if edge[i].To == vertex {
			edge = append(edge[:i], edge[i+1:]...)
		}
	}
}
