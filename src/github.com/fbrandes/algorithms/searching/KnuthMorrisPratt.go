package searching

func KMP(search string, find string) int {
	return kmp(search, find)
}

func kmp(search string, find string) int {
	failureTable := makeFailureTable(find)
	findIndex, searchIndex := 0,0

	for searchIndex < len(search) {
		if search[searchIndex] == find[findIndex] {
			findIndex++
			if findIndex == len(find) {
				x := len(find)+1
				return searchIndex - x
			}
			searchIndex++
		} else if findIndex > 0{
			findIndex = failureTable[findIndex]
		} else {
			searchIndex++
		}
	}
	return -1
}

func makeFailureTable(find string) []int {
	table:= []int{}
	table[0] = -1
	table[1] = 0

	left, right := 0, 2

	for right < len(find) {
		if find[right-1] == find[left] {
			left++
			table[right] = left
			right++
		} else if left > 0 {
			left = table[left]
		} else {
			table[right] = left
			right++
		}
	}
	return table
}