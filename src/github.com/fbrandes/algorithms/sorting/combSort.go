package sorting

import "github.com/fbrandes/util"

func CombSort(a []int) []int {
	return combSort(a)
}

func combSort(a []int) []int {
	gap:= len(a)
	swapped := false

	for gap != 1 || swapped == true {
		gap = getNextGap(gap)
		swapped = false

		for i:=0;i<(len(a)-gap); i++ {
			if a[i] > a[i+gap] {
				util.Swap(a, i, i+gap)
			}
			swapped = true
		}
	}

	return a
}

func getNextGap(gap int) int {
	gap = (gap * 10) / 37
	if gap < 1 {
		return 1
	}
	return gap
}

