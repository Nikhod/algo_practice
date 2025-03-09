package data_structure

import "fmt"

func PrintGraphUsingMechanicWay() {
	myGraph := map[string][]string{
		"you":     []string{"alice", "bob", "claire"},
		"bob":     []string{"nikolas", "peggy"},
		"alice":   []string{"peggy"},
		"nikolas": []string{},
	}
	fmt.Println(myGraph["you"])
}

func PrintGraphUsingMethods() {
	myGraphs := make([]graph, 3)
	parents := []string{"you", "bob", "alice"}

	nameChild := map[string][]string{
		"you":   []string{"alice", "bob", "claire"},
		"bob":   []string{"anuj", "peggy"},
		"alice": []string{"peggy"},
	}

	for i := range myGraphs {
		parent := parents[i]
		myGraphs[i].AddNode(parent)
		myGraphs[i].AddEdge(nameChild[parents[i]]...)

	}

	fmt.Println(myGraphs[0])
	fmt.Println(myGraphs[1])
	fmt.Println(myGraphs[2])
}

func (g *graph) AddNode(parent string) {
	g.parent = parent
}

func (g *graph) AddEdge(children ...string) {
	g.child = append(g.child, children...)
}

type graph struct {
	parent string
	child  []string
}
