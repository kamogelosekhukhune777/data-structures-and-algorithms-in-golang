package unweighted

//graph represents an unweighted graph
// using adjacency list
type Graph struct {
	vertices map[string][]string
}

//NewGraph creates an instance of Graph
func NewGraph() *Graph {
	return &Graph{
		vertices: make(map[string][]string),
	}
}

//AddVertex adds a vertex to the graph
func (g *Graph) AddVertex(vertex string) {
	if _, exists := g.vertices[vertex]; !exists {
		g.vertices[vertex] = []string{}
	}
}
