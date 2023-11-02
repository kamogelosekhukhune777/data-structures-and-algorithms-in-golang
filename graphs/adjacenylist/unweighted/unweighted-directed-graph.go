package unweighted

//AddDirectedEdge adds a directed edge from source to destination
func (g *Graph) AddDirectedEdge(source, destination string) {
	//if the suppossed source vertex does not exist we create it
	if _, exists := g.vertices[source]; !exists {
		g.AddVertex(source)
	}
	//if the suppossed destination vertex does not exist we create it
	if _, exists := g.vertices[destination]; !exists {
		g.AddVertex(destination)
	}
	g.vertices[source] = append(g.vertices[source], destination)
}

//RemoveDirectedEdge removes a directed edge from
// source to destination vertex
func (g *Graph) RemoveDirectedEdge(source, destination string) {
	if _, exists := g.vertices[source]; exists {
		for i, v := range g.vertices[source] {
			if v == destination {
				g.vertices[source] = append(g.vertices[source][:i], g.vertices[source][i+1:]...)
				break
			}
		}
	}
}

//RemoveUnweightedVertex removes a Vertex and all its
// associated edges from the edge
func (g *Graph) RemoveVertex(vertex string) {
	delete(g.vertices, vertex)
	for v := range g.vertices {
		g.RemoveDirectedEdge(v, vertex)
	}
}
