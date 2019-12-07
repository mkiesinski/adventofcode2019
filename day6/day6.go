package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func readFile(filename string) string {
	bytes, err := ioutil.ReadFile(filename)
	check(err)
	return strings.TrimSpace(string(bytes))
}

type node struct {
	id       string
	parent   *node
	parentID string
	children []*node
}

func (n *node) addChild(child *node) {
	n.children = append(n.children, child)
	child.parent = n
}

func (n *node) findID(id string) (*node, bool) {
	if n.id == id {
		return n, true
	}
	for i := range n.children {
		find, found := n.children[i].findID(id)
		if found {
			return find, found
		}
	}
	return nil, false
}

func (n *node) getDepth() int {
	depth := 0
	current := n
	for current.parent != nil {
		depth++
		current = current.parent
	}

	return depth
}

func (n *node) getTotalOrbits() int {
	totalOrbits := n.getDepth()

	for i := range n.children {
		totalOrbits += n.children[i].getTotalOrbits()
	}

	return totalOrbits
}

func (n *node) printNode() {
	fmt.Println(n)
	for i := range n.children {
		n.children[i].printNode()
	}
}

func (n *node) findPathLength(endID string) int {
	length := 0
	current := n

	for current.parent != nil {
		endNode, found := current.findID(endID)
		if found {
			return endNode.getDepth() - current.getDepth() + length - 2
		}
		current = current.parent
		length++
	}

	return -1
}

type tree struct {
	root *node
}

func (t *tree) findID(id string) (*node, bool) {
	find, found := t.root.findID(id)
	return find, found
}

func (t *tree) getTotalOrbits() int {
	return t.root.getTotalOrbits()
}

func (t *tree) insertNode(newNode *node) {
	parent, found := t.findID(newNode.parentID)

	if found {
		parent.addChild(newNode)
	} else {
		t.root.addChild(newNode)
	}

	for i := 0; i < len(t.root.children); i++ {
		parentless := t.root.children[i]
		if parentless.parentID == newNode.id {
			newNode.addChild(parentless)
			t.root.children = append(t.root.children[:i], t.root.children[i+1:]...)
			i--
		}
	}
}

func (t *tree) findPathLength(startID string, endID string) int {
	startNode, foundStart := t.findID(startID)
	_, foundEnd := t.findID(endID)

	if foundStart && foundEnd {
		return startNode.findPathLength(endID)
	}
	return -1
}

func makeTree() tree {
	rootNode := node{id: "COM"}
	return tree{root: &rootNode}
}

func main() {
	orbitTree := makeTree()
	fileString := readFile("input")
	inputArray := strings.Split(strings.Replace(string(fileString), "\r\n", "\n", -1), "\n")

	for i := range inputArray {
		orbitInput := strings.Split(string(inputArray[i]), ")")
		orbitTree.insertNode(&node{id: orbitInput[1], parentID: orbitInput[0]})
	}

	fmt.Println("=== Part 1 ===")
	fmt.Printf("Total of direct and inderect orbits: %d\n", orbitTree.getTotalOrbits())

	fmt.Println("=== Part 2 ===")
	fmt.Printf("Minimum orbital transfers required for YOU to reach SAN: %d\n", orbitTree.findPathLength("YOU", "SAN"))
}
