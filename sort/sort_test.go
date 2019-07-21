package sort

import (
	"fmt"
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	array := []int{3, 5, 2, 6, 1, 4}
	ans := []int{1, 2, 3, 4, 5, 6}
	QuickSort(array)
	fmt.Println(array)
	if !reflect.DeepEqual(array, ans) {
		t.Error("Something wrong with QuickSort")
	}
}

func TestMergeSort(t *testing.T) {
	array := []int{3, 5, 2, 6, 1, 4}
	ans := []int{1, 2, 3, 4, 5, 6}
	array = MergeSort(array)
	fmt.Println(array)
	if !reflect.DeepEqual(array, ans) {
		t.Error("Something wrong with QuickSort")
	}
}

func TestBubbleSort(t *testing.T) {
	array := []int{3, 5, 2, 6, 1, 4}
	ans := []int{1, 2, 3, 4, 5, 6}
	BubbleSort(array)
	fmt.Println(array)
	if !reflect.DeepEqual(array, ans) {
		t.Error("Something wrong with BubbleSort")
	}
}
