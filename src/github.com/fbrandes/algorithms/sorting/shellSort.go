package sorting

func ShellSort(a []int) []int {
	return shellSort(a)
}
func shellSort(a []int) []int {
	cols := []int{1391376, 463792, 198768, 86961, 33936,
		13776, 4592, 1968, 861, 336, 112, 48, 21, 7, 3, 1}

	for k:=0;k<16;k++ {
		h:=cols[k]
		for i:=h; i< len(a); i++ {
			j:=i
			t :=a[j]
			for j>=h && a[j-h]> t {
				a[j] = a[j-h]
				j -= h
			}
			a[j]=t
		}
	}
	return a
}
