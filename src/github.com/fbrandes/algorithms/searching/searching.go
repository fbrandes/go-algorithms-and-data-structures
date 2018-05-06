package searching

func LinearSearch(a []int, n int) int {
	for i:=0;i<=len(a);i++ {
		if n == a[i] {
			return n
		}
	}
	return -1
}

func BinarySearch(a []int, key int) bool {
	return binarySearch(a, key, 0, len(a))
}

// TODO implement with slices to avoid copying the whole array
func binarySearch(a []int, key, left, right int) bool {
	if right <= left {
		return false
	}
	pivot := left + (right - left) / 2

	if pivot > key {
		return binarySearch(a, key, left, pivot)
	} else if pivot < key {
		return binarySearch(a, key, pivot+1, right)
	} else {
		return true
	}
}
