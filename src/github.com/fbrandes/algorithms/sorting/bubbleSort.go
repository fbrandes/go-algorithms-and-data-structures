package sorting

func BubbleSort(a []int) []int {
	n := len(a)
	for j := n; j > 0; j-- {
		for i := 1; i <= n-1; i++ {
			if a[i-1] > a[i] {
				swap(a, i-1, i)
			}
		}
	}
	return a
}
