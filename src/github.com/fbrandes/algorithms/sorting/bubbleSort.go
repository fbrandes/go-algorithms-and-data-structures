package sorting

import "github.com/fbrandes/util"

func BubbleSort(a []int) []int {
	n := len(a)
	for j := n; j > 0; j-- {
		for i := 1; i <= n-1; i++ {
			if a[i-1] > a[i] {
				util.Swap(a, i-1, i)
			}
		}
	}
	return a
}
