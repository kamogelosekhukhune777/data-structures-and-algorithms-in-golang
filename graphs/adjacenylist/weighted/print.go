package weighted

import "fmt"

func (g *Graph) PrintGraph() {
	for vertex, edges := range g.adjacencyList {
		fmt.Printf("vertex %d ->", vertex)
		for _, edge := range edges {
			fmt.Printf("(%d, %d)", edge.To, edge.Weight)
		}
		fmt.Println()
	}

}
