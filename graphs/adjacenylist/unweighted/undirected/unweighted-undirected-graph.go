package unweighted

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
func (g *Graph) RemoveVertex(vertex string) {
	adjacentVertices := g.vertices[vertex]
	delete(g.vertices, vertex)
	for _, v := range adjacentVertices {
		g.RemoveUndirectedEdge(vertex, v)
	}
}

//PrintGraph prints the graph's adjacency
// list representation
func (g *Graph) PrintGraph() {
	for vertex, edges := range g.vertices {
		fmt.Printf("%s -> %s\n", vertex, edges)
	}
}
