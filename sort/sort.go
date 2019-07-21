package sort

func QuickSort(source []int) {
	if len(source) < 2 {
		return
	}
	pivot := source[0]
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

func MergeSort(source []int) []int {
	if len(source) < 2 {
		return source
	}
	m := int(len(source) / 2)
	var left []int = MergeSort(source[:m])
	var right []int = MergeSort(source[m:])
	dist := make([]int, len(source))
	var i, j, k int = 0, 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			dist[k] = left[i]
			i++
		} else {
			dist[k] = right[j]
			j++
		}
		k++
	}
	if i < len(left) {
		for i < len(left) {
			dist[k] = left[i]
			i++
			k++
		}
	} else {
		for j < len(right) {
			dist[k] = right[j]
			j++
			k++
		}
	}
	return dist
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
