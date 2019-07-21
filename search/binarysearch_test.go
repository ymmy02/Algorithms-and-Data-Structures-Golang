package binarysearch

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	primeArray := []int{2, 3, 5, 7, 11, 13, 17, 19, 23}

	if BinarySearch(primeArray, 5) != 2 {
		t.Error("Something wrong with BinarySearch")
	}
	if BinarySearch(primeArray, 17) != 6 {
		t.Error("Something wrong with BinarySearch")
	}
	if BinarySearch(primeArray, 100) != -1 {
		t.Error("Something wrong with BinarySearch")
	}
}
