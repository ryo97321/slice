package slice

// スライスaからnumberを削除します
func Remove(a []int, number int) []int {
	var index int
	for i := 0; i < len(a); i++ {
		if a[i] == number {
			index = i
		}
	}

	result := a[:index]
	if index+1 < len(a) {
		for i := index + 1; i < len(a); i++ {
			result = append(result, a[i])
		}
	}

	return result
}
