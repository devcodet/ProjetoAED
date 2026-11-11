package main_stack

type Node struct {
	word string
	next *Node
}

type Pilha struct {
	topo *Node
}
