package graphs_test

import (
	"testing"

	"github.com/devlipe/data_structures/graphs"
	st "github.com/devlipe/data_structures/structures"
)

func TestFloydWarshall(t *testing.T) {
	graph := st.NewDigraph[string]()

	graph.AddEdge("a", "b", 1)
	graph.AddEdge("a", "c", 3)
	graph.AddEdge("b", "g", 5)
	graph.AddEdge("c", "g", 8)
	graph.AddEdge("g", "h", 6)
	graph.AddEdge("c", "d", 2)
	graph.AddEdge("g", "f", 4)
	graph.AddEdge("d", "f", 3)
	graph.AddEdge("d", "e", 5)

	m := graphs.FloydWarshall(graph)
	if c := m["a"]["e"]; c != 10 {
		t.Errorf("bad shortest cost %f for a-e", c)
	}
}
