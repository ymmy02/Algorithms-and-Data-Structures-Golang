package binarysearch

func BinarySearch(array []int, value int) int {
	var l int = 0
	var h int = len(array) - 1
	var m int = int((l + h) / 2)

	for l < h-1 {
		if array[m] > value {
			h = m
		} else {
			l = m
		}
		m = int((l + h) / 2)
	}
	if array[l] == value {
		return l
	} else {
		return -1
	}
}
