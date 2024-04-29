package swap

func SwapContents(listObj []int) {
	for i, j := 0, len(listObj)-1; i < j; i, j = i+1, j-1 {
		listObj[i], listObj[j] = listObj[j], listObj[i]
	}
}
