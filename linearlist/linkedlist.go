package linearlist

import (
	"fmt"
)

type Node struct {
	value int
	next  *Node
}

type LinkedList struct {
	head *Node
}

func (linkedList *LinkedList) Append(value int) {
	var lastNode *Node
	lastNode = linkedList.head
	newNode := &Node{value: value, next: nil}
	if lastNode == nil {
		linkedList.head = newNode
	} else {
		for lastNode.next != nil {
			lastNode = lastNode.next
		}
		lastNode.next = newNode
	}
}

func (linkedList *LinkedList) PrintValues() {
	var currentNode = linkedList.head
	for currentNode != nil {
		fmt.Printf("%d ", currentNode.value)
		currentNode = currentNode.next
	}
	fmt.Printf("\n")
}
