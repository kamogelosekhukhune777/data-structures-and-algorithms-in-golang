package weighted

import "fmt"

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

func (g *Graph) AddEdge(from, to, weight int) {
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

func (g *Graph) RemoveEdge(from, to int) {
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

func (g *Graph) RemoveVertex(vertex int) {
	//Remove vertex & its edges from the adjacency list
	delete(g.adjacencyList, vertex)
	//Remove all edges pointing to the removed vertex
	for i, edge := range g.adjacencyList[vertex] {
		if edge.To == vertex {
			edge = append(edge[:i], edge[i+1:]...)
		}
	}
}

func (g *Graph) PrintGraph() {
	for vertex, edges := range g.adjacencyList {
		fmt.Printf("vertex %d ->", vertex)
		for _, edge := range edges {
			fmt.Printf("(%d, %d)", edge.To, edge.Weight)
		}
		fmt.Println()
	}

}
