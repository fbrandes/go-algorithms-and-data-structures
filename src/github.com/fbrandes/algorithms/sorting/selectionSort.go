package sorting

import "github.com/fbrandes/util"

func SelectionSort(a []int) []int {
	return selectionSort(a)
}

func selectionSort(a []int) []int {
	for i:=0 ; i < len(a); i++ {
		min := i
		for j:=i+1; j < len(a); j++ {
			if a[j] < a[min] {
				min = j
			}
		}
		util.Swap(a, i, min)
	}
	return a
}
