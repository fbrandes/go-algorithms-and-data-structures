package sorting

func swap(a []int, i int, j int) {
	temp := a[j]
	a[j] = a[i]
	a[i] = temp
}

func getMax(a []int) int {
	max := 0
	for i := 0; i < len(a); i++ {
		if a[i] > max {
			max = a[i]
		}
	}
	return max
}
