package main

type Node struct {
	word     string
	nextNode *Node
}

type Pilha struct {
	topNode *Node
}

func (p *Pilha) Push(word string) {
	p.topNode = &Node{word: word, nextNode: p.topNode}
}

func (p *Pilha) Pop() (string, bool) {
	if p.topNode == nil {
		return "", false
	}
	word := p.topNode.word
	p.topNode = p.topNode.nextNode
	return word, true
}

func (p *Pilha) IsEmpty() bool {
	return p.topNode == nil
}
