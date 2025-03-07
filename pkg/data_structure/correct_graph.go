package data_structure

import "fmt"

type Node struct {
	Name      string
	Neighbors []Node
}

func CorrectPrintingGraph() {
	you := newParent("you")
	bob := newParent("bob")
	alice := newParent("alice")
	claire := newParent("claire")
	nikolas := newParent("nikolas")
	peggy := newParent("peggy")
	thom := newParent("thom")
	jonny := newParent("jonny")

	alice.AddNeighbor(peggy)
	claire.AddNeighbor(thom)
	claire.AddNeighbor(jonny)

	bob.AddNeighbor(nikolas)
	bob.AddNeighbor(peggy)

	you.AddNeighbor(bob)
	you.AddNeighbor(alice)
	you.AddNeighbor(claire)

	//fmt.Printf("you: %+v\nclaire: %+v\nbob: %+v\n", you, claire, bob)
	for i, value := range you.Neighbors {
		fmt.Printf("%v neighbor: %+v\n", i+1, value)
	}
}

func newParent(parentName string) *Node {
	return &Node{Name: parentName}
}

func (n *Node) AddNeighbor(neighbor *Node) {
	n.Neighbors = append(n.Neighbors, *neighbor)

}
