package graph

import (
	"math"
	"testing"
)

func Test_NewGraph(t *testing.T) {
	g := NewGraph(3)
	if g.vertices != 3 {
		t.Errorf("expected 3 vertices, got %d", g.vertices)
	}
	if len(g.edges) != 0 {
		t.Errorf("expected 0 edges, got %d", len(g.edges))
	}
}

func Test_AddEdge(t *testing.T) {
	g := NewGraph(3)
	g.AddEdge(0, 1, 0.5)
	g.AddEdge(1, 2, 2.0)

	if len(g.edges) != 2 {
		t.Errorf("expected 2 edges, got %d", len(g.edges))
	}

	expectedWeight := -math.Log(0.5)
	if g.edges[0].weight != expectedWeight {
		t.Errorf("expected weight %f, got %f", expectedWeight, g.edges[0].weight)
	}
}

func Test_DetectArbitrage(t *testing.T) {
	g := NewGraph(4)
	g.AddEdge(0, 1, 0.5)
	g.AddEdge(1, 2, 2.0)
	g.AddEdge(2, 3, 1.5)
	g.AddEdge(3, 0, 0.75)

	if !g.DetectArbitrage(0) {
		t.Error("expected arbitrage opportunity to be detected, but it was not")
	}

	g = NewGraph(3)
	g.AddEdge(0, 1, 1.0)
	g.AddEdge(1, 2, 1.0)
	g.AddEdge(2, 0, 1.0)

	if g.DetectArbitrage(0) {
		t.Error("expected no arbitrage opportunity to be detected, but it was")
	}
}
