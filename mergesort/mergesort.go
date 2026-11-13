package mergesort

func mergeSort(words []string) []string {
	if len(words) <= 1 {
		return words
	}

	half := len(words) / 2
	left := mergeSort(words[:half])
	right := mergeSort(words[half:])

	return merge(left, right)
}

func merge(left, right []string) []string {
	result := make([]string, 0, len(left)+len(right))
	i := 0
	j := 0

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
