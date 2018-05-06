package sorting

const BASE = 10
const SIGNIFICANTDIGIT = 1

func RadixSort(a []int) []int {
	neg, pos := divide(a)
	return append(radixSort(neg), radixSort(pos)...)
}

func divide(a []int) ([]int, []int) {
	var pos, neg []int
	for i := 0; i < len(a); i++ {
		if a[i] < 0 {
			neg = append(neg, a[i])
		} else {
			pos = append(pos, a[i])
		}
	}
	return neg, pos
}

func radixSort(a []int) []int {
	largestNum := getMax(a)
	significantDigit := SIGNIFICANTDIGIT
	helper := make([]int, len(a), len(a))

	for largestNum/significantDigit > 0 {
		bucket := [BASE]int{0}
		for i := 0; i < len(a); i++ {
			bucket[(a[i]/significantDigit)%BASE]++
		}
		for i := 1; i < BASE; i++ {
			bucket[i] += bucket[i-1]
		}
		for i := len(a) - 1; i >= 0; i-- {
			bucket[(a[i]/significantDigit)%BASE]--
			helper[bucket[(a[i]/significantDigit)%BASE]] = a[i]
		}
		for i := 0; i < len(a); i++ {
			a[i] = helper[i]
		}

		significantDigit *= BASE
	}
	return a
}
