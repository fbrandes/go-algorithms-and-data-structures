package sorting

func InsertionSort(a []int) []int {
	return insertionSort(a)
}

func insertionSort(a []int) []int {
	for i := 1; i < len(a); i++ {
		temp := a[i]
		j:=i
		for j > 0 && a[j-1]>temp {
			a[j] = a[j-1]
			j--
		}
		a[j] = temp
	}
	return a
}
