package pilha_slice

type PilhaSlice struct {
	items []string
}

func (p *PilhaSlice) Push(word string) {
	p.items = append(p.items, word)
}

func (p *PilhaSlice) Pop() (string, bool) {
	if p.IsEmpty() {
		return "", false
	}
	last := len(p.items) - 1
	word := p.items[last]
	p.items = p.items[:last]
	return word, true
}

func (p *PilhaSlice) IsEmpty() bool {
	return len(p.items) == 0
}

func (p *PilhaSlice) Empty() bool {
	p.items = nil
	return true
}
