package shuffling

import (
	"math/rand"
	"github.com/fbrandes/util"
)

func FisherYatesShuffle(a []int) []int {
	ret := a
	return fisherYatesShuffle(ret)
}

func fisherYatesShuffle(a []int) []int {
	for i := len(a)-1; i > 0; i-- {
		index := rand.Intn(i + 1)
		util.Swap(a, index, i)
	}
	return a
}
