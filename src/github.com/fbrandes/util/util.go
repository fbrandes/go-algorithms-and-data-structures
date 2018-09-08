package util

func Swap(a []int, i int, j int) {
	temp := a[j]
	a[j] = a[i]
	a[i] = temp
}

func GetMax(a []int) int {
	max := a[0]
	for i := 0; i < len(a); i++ {
		if a[i] > max {
			max = a[i]
		}
	}
	return max
}

func GetMin(a []int) int {
	min := a[0]
	for i := 0; i < len(a); i++ {
		if a[i] < min {
			min = a[i]
		}
	}
	return min
}

//func GetMin(a, b int) int {
//	if a != b {
//		if a < b {
//			return a
//		} else {
//			return b
//		}
//	}
//	return -1
//}
