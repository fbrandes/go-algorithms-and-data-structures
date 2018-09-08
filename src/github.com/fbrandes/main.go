package main

import (
	//"./datastructures/basic"
	"fmt"
	"github.com/fbrandes/algorithms/shuffling"
	"github.com/fbrandes/algorithms/sorting"
)

func main() {
	b := []int{9, 3, 25, 6, 87, 4, 36, 2}
	c := sorting.CountingSort(b)
	fmt.Println(c)

	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	a = shuffling.FisherYatesShuffle(a)
	fmt.Println(sorting.CombSort(a))

}
