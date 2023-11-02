package adjacenylist

import "fmt"

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

//AddEdge adds a directed edge from source to destination
func (g *Graph) AddEdge(source, destination string) {
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

//RemoveEdge removes a directed edge from
// source to destination vertex
func (g *Graph) RemoveEdge(source, destination string) {
	if _, exists := g.vertices[source]; exists {
		for i, v := range g.vertices[source] {
			if v == destination {
				g.vertices[source] = append(g.vertices[source][:i], g.vertices[source][i+1:]...)
				break
			}
		}
	}
}

//RemoveVertex removes a Vertex and all its
// associated edges from the edge
func (g *Graph) RemoveVertex(vertex string) {
	delete(g.vertices, vertex)
	for v := range g.vertices {
		g.RemoveEdge(v, vertex)
	}
}

func (g *Graph) PrintGraph() {
	for vertex, edges := range g.vertices {
		fmt.Printf("%s -> %s\n", vertex, edges)
	}
}
