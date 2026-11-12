package pilha_slice

type Pilha struct {
	items []string
}

func (p *Pilha) Push(word string) {
	p.items = append(p.items, word)
}

func (p *Pilha) Pop() (string, bool) {
	if len(p.items) == 0 {
		return "", false
	}
	last := len(p.items) - 1
	word := p.items[last]
	p.items = p.items[:last]
	return word, true
}

func (p *Pilha) IsEmpty() bool {
	return len(p.items) == 0
}
