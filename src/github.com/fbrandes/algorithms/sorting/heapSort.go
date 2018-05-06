package sorting

func HeapSort(a []int) []int {
	return heapSort(a)
}

func heapSort(a []int) []int {
	makeHeap(a)

	for i := len(a)-1; i > 0; i-- {
		swap(a, i, 0)
		downHeap(a, 0, i)
	}
	return a
}

func makeHeap(a []int) {
	for i:=(len(a)/2)-1; i >= 0; i-- {
		downHeap(a, i, len(a))
	}
}

func downHeap(a []int, i,n int) {
	for i <= (n/2)-1 {
		childIndex := ((i+1)*2)-1

		if childIndex+1 <= n-1 {
			if a[childIndex] < a[childIndex+1] {
				childIndex++
			}
		}
		if a[i] < a[childIndex] {
			swap(a, i, childIndex)
			i = childIndex
		} else {
			break
		}
	}
}