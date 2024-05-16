package main

import (
	"github.com/mdshahjahanmiah/trading-arbitrage-detector/pkg/graph"
	"log/slog"
)

func main() {
	// A new graph is created with 4 vertices. Each vertex represents a different currency.
	graph := graph.NewGraph(4)

	// Representing exchange rates with negative logarithms. Exchange rates are converted to
	// negative logarithms to transform the problem into the shortest path problem.
	graph.AddEdge(0, 1, 0.5)  // USD to EUR
	graph.AddEdge(1, 2, 2.0)  // EUR to GBP
	graph.AddEdge(2, 3, 1.5)  // GBP to JPY
	graph.AddEdge(3, 0, 0.75) // JPY to USD

	// The method is called with source vertex 0, representing the starting currency (USD in this case).
	// currencyIndex: [USD": 0,"EUR": 1, "GBP": 2, "JPY": 3]
	if graph.DetectArbitrage(0) {
		slog.Info("Arbitrage opportunity detected!")
	} else {
		slog.Info("No arbitrage opportunity found.")
	}
}
