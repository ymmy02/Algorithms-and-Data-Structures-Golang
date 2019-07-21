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

func BubbleSort(source []int) {
	var length int = len(source)
	if length < 2 {
		return
	}

	for i := 0; i < length-1; i++ {
		var swap bool = false
		for j := length - 1; j > i; j-- {
			if source[j] < source[j-1] {
				source[j], source[j-1] = source[j-1], source[j]
				swap = true
			}
		}
		if !swap {
			break
		}
	}
}
