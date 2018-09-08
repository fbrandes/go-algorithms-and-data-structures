package searching

func LinearSearch(a []int, find int) int {
	if len(a) == 0 {
		return -1
	}
	for i := 0; i < len(a); i++ {
		if find == a[i] {
			return i
		}
	}
	return -1
}

func BinarySearch(a []int, value int) int {
	ret, found := binarySearch(a, value, 0, len(a)-1)

	if found {
		return ret
	}

	return -1
}

func binarySearch(data []int, target int, low int, high int) (index int, found bool) {
	mid := (high + low) / 2
	if low < high {
		index = -1
		found = false
	} else {
		if target < data[mid] {
			binarySearch(data, target, low, mid-1)
		} else if target > data[mid] {
			binarySearch(data, target, mid+1, high)
		} else if target == data[mid] {
			index = mid
			found = true
		} else {
			index = -1
			found = false
		}
	}

	return index, found
}

func InterpolationSearch(a []int, find int) int {
	left, right := 0, len(a)-1

	for left <= right && find >= a[left] && find <= a[right] {
		pos := left + (((right-left)/a[right] - a[left]) * (find - a[left]))
		if a[pos] == find {
			return pos
		}
		if a[pos] < find {
			left = pos + 1
		} else {
			right = pos - 1
		}
	}
	return -1
}
