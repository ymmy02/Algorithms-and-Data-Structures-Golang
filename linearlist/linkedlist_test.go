package linearlist

import (
	"testing"
)

func TestPrintValues(t *testing.T) {
	linkedList := LinkedList{head: nil}
	linkedList.Append(1)
	linkedList.Append(2)
	linkedList.Append(3)
	linkedList.Append(1)
	linkedList.Append(5)
	if linkedList.ToString() != "1 2 3 1 5 " {
		t.Error("Something worng with Append or ToString method")
	}
	linkedList.PrintValues()
}

func TestDeleteOne(t *testing.T) {
	linkedList := LinkedList{head: nil}
	linkedList.Append(1)
	linkedList.Append(2)
	linkedList.Append(3)
	linkedList.Append(1)
	linkedList.Append(5)
	linkedList.DeleteOne(100)
	if linkedList.ToString() != "1 2 3 1 5 " {
		t.Error("Something worng with DeleteOne method")
	}
	linkedList.DeleteOne(1)
	linkedList.DeleteOne(5)
	if linkedList.ToString() != "2 3 1 " {
		t.Error("Something worng with DeleteOne method")
	}
}

func TestReverse(t *testing.T) {
	linkedList := LinkedList{head: nil}
	linkedList.Append(1)
	linkedList.Append(2)
	linkedList.Append(3)
	linkedList.Append(1)
	linkedList.Append(5)

	linkedList.Reverse()
	if linkedList.ToString() != "5 1 3 2 1 " {
		t.Error("Something worng with Reverse method")
	}
}
