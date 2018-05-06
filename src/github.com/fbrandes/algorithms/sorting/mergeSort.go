package sorting

func MergeSort(a []int) []int {
	if len(a) < 2 {
		return a
	}
	pivot := len(a) / 2
	return merge(MergeSort(a[:pivot]), MergeSort(a[pivot:]), len(a))
}

func merge(left, right []int, size int) []int{
	i, j := 0, 0
	arr := make([]int, size, size)

	for k := 0; k < size; k++ {
		if i > len(left)-1 && j <= len(right)-1 {
			arr[k] = right[j]
			j++
		} else if j > len(right)-1 && i <= len(left)-1 {
			arr[k] = left[i]
			i++
		} else if left[i] < right[j] {
			arr[k] = left[i]
			i++
		} else {
			arr[k] = right[j]
			j++
		}
	}
	return arr
}
