package pilha_list

type Node struct {
	word     string
	nextNode *Node
}

type PilhaList struct {
	topNode *Node
}

func (p *PilhaList) Push(word string) {
	p.topNode = &Node{word: word, nextNode: p.topNode}
}

func (p *PilhaList) Pop() (string, bool) {
	if p.topNode == nil {
		return "", false
	}
	word := p.topNode.word
	p.topNode = p.topNode.nextNode
	return word, true
}

func (p *PilhaList) IsEmpty() bool {
	return p.topNode == nil
}

func (p *PilhaList) Empty() bool {
	p.topNode = nil
	return true
}
