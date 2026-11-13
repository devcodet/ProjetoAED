package main

func insertionSort(words []string) {
	for i := 1; i < len(words); i++ {
		key := words[i]
		j := i - 1

		for j >= 0 && words[j] > key {
			words[j+1] = words[j]
			j--
		}
		words[j+1] = key
	}
}
