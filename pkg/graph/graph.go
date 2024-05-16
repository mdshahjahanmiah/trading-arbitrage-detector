package graph

import (
	"math"
)

type Edge struct {
	from, to int
	weight   float64
}

type Graph struct {
	vertices int
	edges    []Edge
}

// NewGraph creates a new graph with a specified number of vertices
func NewGraph(vertices int) *Graph {
	return &Graph{vertices: vertices, edges: []Edge{}}
}

// AddEdge adds a directed edge with a weight (negative log of the rate) between two vertices
func (g *Graph) AddEdge(from, to int, rate float64) {
	weight := -math.Log(rate)
	g.edges = append(g.edges, Edge{from: from, to: to, weight: weight})
}

// In trading, arbitrage refers to the practice of taking advantage of a price difference between two or more markets to generate a profit.
// For example, if a currency can be exchanged for another currency at a lower rate in one market and sold at a higher rate in another,
// an arbitrage opportunity exists. By modeling the currencies and exchange rates as a graph, where vertices represent currencies and
// edges represent exchange rates, we can detect these opportunities by identifying negative weight cycles in the graph.

// A negative weight cycle occurs when the sum of the weights of the edges in a cycle is negative, indicating that there is a way to
// repeatedly traverse the cycle and decrease the total weight indefinitely. In the context of currency exchange, this means that
// we can keep exchanging currencies in a cycle and end up with more of the starting currency than we began with, essentially making a profit.

// The Bellman-Ford algorithm is particularly suited for this task because it can handle graphs with negative weights and can detect
// negative weight cycles.

// DetectArbitrage uses the Bellman-Ford algorithm to detect negative weight cycles, which indicate arbitrage opportunities in trading.
func (g *Graph) DetectArbitrage(src int) bool {
	// Step 1: Initialize distances from the source to all other vertices as infinity, except the source itself.
	// This represents the initial assumption that all other currencies are infinitely far from the source currency.
	distances := make([]float64, g.vertices)
	for i := range distances {
		distances[i] = math.Inf(1)
	}
	distances[src] = 0

	// Step 2: Relax edges repeatedly.
	// Relaxing an edge means trying to improve the shortest known distance to a vertex by using an edge from another vertex.
	// If a shorter path is found, the distance to the vertex is updated.
	for i := 1; i < g.vertices; i++ {
		for _, edge := range g.edges {
			if distances[edge.from] != math.Inf(1) && distances[edge.from]+edge.weight < distances[edge.to] {
				distances[edge.to] = distances[edge.from] + edge.weight
			}
		}
	}

	// Step 3: Check for negative weight cycles.
	// After relaxing all edges, we do one more pass over all edges to see if we can still find a shorter path.
	// If we can, it means there is a negative weight cycle in the graph, indicating an arbitrage opportunity.
	for _, edge := range g.edges {
		if distances[edge.from] != math.Inf(1) && distances[edge.from]+edge.weight < distances[edge.to] {
			return true // Negative cycle detected, indicating an arbitrage opportunity
		}
	}

	// No negative cycles detected, hence no arbitrage opportunities.
	return false
}
