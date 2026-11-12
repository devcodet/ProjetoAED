package pilha_list

type Node struct {
	word     string
	nextNode *Node
}

type PilhaStack struct {
	topNode *Node
}

func (p *PilhaStack) Push(word string) {
	p.topNode = &Node{word: word, nextNode: p.topNode}
}

func (p *PilhaStack) Pop() (string, bool) {
	if p.topNode == nil {
		return "", false
	}
	word := p.topNode.word
	p.topNode = p.topNode.nextNode
	return word, true
}

func (p *PilhaStack) IsEmpty() bool {
	return p.topNode == nil
}
