package main

func selectionSort(words []string) {
	for i := 0; i < len(words)-1; i++ {
		min := i
		for j := i + 1; j < len(words); j++ {
			if words[j] < words[min] {
				min = j
			}
		}
		if min != i {
			words[i], words[min] = words[min], words[i]
		}
	}
}
