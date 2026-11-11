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

func mergeSort(words []string) []string {
	if len(words) <= 1 {
		return words
	}

	meio := len(words) / 2
	left := mergeSort(words[:meio])
	right := mergeSort(words[meio:])

	return merge(left, right)
}

func merge(left, right []string) []string {
	result := make([]string, 0, len(left)+len(right))
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}
