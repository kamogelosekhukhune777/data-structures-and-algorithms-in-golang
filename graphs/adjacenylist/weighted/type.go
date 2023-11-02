package weighted

type Edge struct {
	To     int
	Weight int
}

// graph represents an unweighted graph
// using adjacency list
type Graph struct {
	adjacencyList map[int][]Edge
}

// NewGraph creates an instance of Graph
func NewGraph() *Graph {
	return &Graph{
		adjacencyList: make(map[int][]Edge),
	}
}

func (g *Graph) AddVertex(vertex int) {
	if _, exists := g.adjacencyList[vertex]; !exists {
		g.adjacencyList[vertex] = []Edge{}
	}
}
