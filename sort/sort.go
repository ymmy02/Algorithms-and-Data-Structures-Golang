package sort

func QuickSort(source []int) {
	if len(source) < 2 {
		return
	}
	m := int(len(source) / 2)
	pivot := source[m]
	l := 0
	r := len(source) - 1
	for {
		for source[l] < pivot {
			l++
		}
		for source[r] > pivot {
			r--
		}
		if l >= r {
			break
		}
		source[l], source[r] = source[r], source[l]
		l++
		r--
	}
	left := source[:l]
	right := source[r+1:]
	QuickSort(left)
	QuickSort(right)
}
