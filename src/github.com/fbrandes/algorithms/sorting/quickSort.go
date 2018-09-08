package sorting

import "github.com/fbrandes/util"

func QuickSort(a []int) []int {
	return quickSort(a, 0, len(a)-1)
}
func quickSort(a []int, left, right int) []int {
	if left <= right {
		pivot := partition(a, left, right)
		quickSort(a, left, pivot-1)
		quickSort(a, pivot+1, right)
	}
	return a
}

func partition(a []int, left, right int) int{
	pivot := a[right]
	i := left -1

	for j:=left;j<=right-1; j++{
		if a[j] < pivot {
			i++
			util.Swap(a, i, j)
		}
	}
	util.Swap(a, i+1, right)
	return i+1
}
