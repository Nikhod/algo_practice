package data_structure

import "fmt"

type Node struct {
	Name      string
	Neighbors []*Node // внезависимости от последовательности заполнения соседей, все равно you (центральный узел)
	// все равно будет иметь доступ (Будет заполнен друзьями, и друзьями их друзей)
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

	you.AddNeighbor(bob)
	you.AddNeighbor(alice)
	you.AddNeighbor(claire)

	alice.AddNeighbor(peggy)
	claire.AddNeighbor(thom)
	claire.AddNeighbor(jonny)

	bob.AddNeighbor(nikolas)
	bob.AddNeighbor(peggy)

	fmt.Println(len(you.Neighbors[0].Neighbors))

}

func newParent(parentName string) *Node {
	return &Node{Name: parentName}
}

func (n *Node) AddNeighbor(neighbor *Node) {
	n.Neighbors = append(n.Neighbors, neighbor)

}
