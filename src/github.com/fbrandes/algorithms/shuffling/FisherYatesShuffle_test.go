package shuffling

import (
	"testing"
)

func TestShuffling(t *testing.T) {
	var results [10][]int

	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

	for i := 0; i < 10; i++ {
		results[i] = FisherYatesShuffle(input)
	}

	if compare(results) == false {
		t.Errorf("Error: Found equal arrays after shuffling")
	}
	t.Logf("Success")
}

func compare(results [10][]int) bool {
	for i := 1; i < len(results); i++ {
		if isEqual(results[i-1], results[i]) {
			return false
		}
	}
	return true
}

func isEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	if len(a) == 0 || len(b) == 0 {
		// at this point both arrays are of equal size, so if either is of length 0 then the other is as well and this means they are equal
		return true
	}

	for i := 0; i < len(a); i++ {
		if !(a[i] == b[i]) {
			return false
		}
	}
	return true
}
