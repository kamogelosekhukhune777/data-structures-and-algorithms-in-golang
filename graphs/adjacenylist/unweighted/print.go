package unweighted

import "fmt"

//PrintGraph prints the graph's adjacency
// list representation
func (g *Graph) PrintGraph() {
	for vertex, edges := range g.vertices {
		fmt.Printf("%s -> %s\n", vertex, edges)
	}
}
