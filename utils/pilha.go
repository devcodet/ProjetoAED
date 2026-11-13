package utils

type Pilha interface {
	Push(word string)
	Pop() (string, bool)
	IsEmpty() bool
	Empty() bool
}
