package sorting

import (
	"github.com/fbrandes/util"
	"runtime"
	"strings"
	"fmt"
)

func CountingSort(a []int) []int {
	return countingSort(a)
}

func countingSort(a []int) []int {
	max := util.GetMax(a)
	min := util.GetMin(a)

	defer func() {
		if x := recover(); x != nil {
			if _, ok := x.(runtime.Error); ok && strings.HasSuffix(x.(error).Error(), "index out of range") {
				fmt.Println("data value out of rangfe (%d..%d)\n", min, max)
				return
			}
		}
	}()

	count := make([]int, max - min + 1)

	for _, x := range count {
		count[x-min]++
	}

	z := 0

	for i, c := range count {
		for ; c > 0; c-- {
			a[z] = i + min
			z++
		}
	}

	return count
}
