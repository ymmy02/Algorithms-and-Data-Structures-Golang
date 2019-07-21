package linearlist

import (
	"fmt"
	"strconv"
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

func (linkedList *LinkedList) DeleteOne(value int) {
	var prevNode, currentNode *Node
	prevNode = nil
	currentNode = linkedList.head
	if currentNode == nil {
		return
	}
	if currentNode.value == value { // Delete the first node
		linkedList.head = currentNode.next
		currentNode = nil
		return
	}
	for currentNode != nil {
		if currentNode.value == value {
			prevNode.next = currentNode.next
			currentNode = nil
			return
		}
		prevNode = currentNode
		currentNode = currentNode.next
	}
}

func (linkedList *LinkedList) Length() int {
	length := int(0)
	var currentNode *Node
	currentNode = linkedList.head
	for currentNode != nil {
		length++
		currentNode = currentNode.next
	}
	return length
}

func (linkedList *LinkedList) PrintValues() {
	var currentNode = linkedList.head
	for currentNode != nil {
		fmt.Printf("%d ", currentNode.value)
		currentNode = currentNode.next
	}
	fmt.Printf("\n")
}

func (linkedList *LinkedList) ToString() string {
	var str string
	var currentNode = linkedList.head
	for currentNode != nil {
		str += strconv.Itoa(currentNode.value) + " "
		currentNode = currentNode.next
	}
	return str
}

func (linkedList *LinkedList) Reverse() {
	if linkedList.head == nil {
		return
	}

	var currentNode, nextNode *Node
	currentNode = linkedList.head
	nextNode = currentNode.next
	currentNode.next = nil
	for nextNode != nil {
		tmpNode := nextNode.next
		nextNode.next = currentNode
		currentNode = nextNode
		nextNode = tmpNode
	}
	linkedList.head = currentNode
}
