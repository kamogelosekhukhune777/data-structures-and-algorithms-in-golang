package unweighted

//AddUndirectedEdge add's an undirected edge between
// two vertices
func (g *Graph) AddUndirectedEdge(vertex1, vertex2 string) {
	if _, exists := g.vertices[vertex1]; !exists {
		g.AddVertex(vertex1)
	}
	if _, exists := g.vertices[vertex2]; !exists {
		g.AddVertex(vertex2)
	}
	g.vertices[vertex1] = append(g.vertices[vertex1], vertex2)
	g.vertices[vertex2] = append(g.vertices[vertex2], vertex1)
}

//RemoveUndirectedEdge removes an undirected edge between
// two vertices
func (g *Graph) RemoveUndirectedEdge(vertex1, vertex2 string) {
	if _, exists1 := g.vertices[vertex1]; exists1 {
		for i, v := range g.vertices[vertex1] {
			if v == vertex2 {
				g.vertices[vertex1] = append(g.vertices[vertex1][:i], g.vertices[vertex1][i+1:]...)
			}
		}
	}
	if _, exists2 := g.vertices[vertex2]; exists2 {
		for i, v := range g.vertices[vertex2] {
			if v == vertex2 {
				g.vertices[vertex2] = append(g.vertices[vertex2][:i], g.vertices[vertex2][i+1:]...)
			}
		}
	}
}

//RemoveUndirectedVertex removes a vertex and all its
// associated edges from the graph
func (g *Graph) RemoveUndirectedVertex(vertex string) {
	adjacentVertices := g.vertices[vertex]
	delete(g.vertices, vertex)
	for _, v := range adjacentVertices {
		g.RemoveUndirectedEdge(vertex, v)
	}
}
