package insertion

func insertionSort(words []string) {
	for i := 1; i < len(words); i++ {
		chave := words[i]
		j := i - 1

		for j >= 0 && words[j] > chave {
			words[j+1] = words[j]
			j--
		}
		words[j+1] = chave
	}
}
