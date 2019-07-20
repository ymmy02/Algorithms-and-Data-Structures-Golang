package linearlist

import (
	"testing"
)

func TestPrintValues(t *testing.T) {
	linkedList := LinkedList{head: nil}
	linkedList.append(1)
	linkedList.append(2)
	linkedList.append(3)
	linkedList.append(1)
	linkedList.append(5)
	linkedList.printValues()
}
