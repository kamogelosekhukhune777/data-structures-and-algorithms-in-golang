package linkedlist

import "fmt"

type DoublyNode struct {
	data     int
	next     *DoublyNode
	previous *DoublyNode
}

type DoublyLinkedList struct {
	head *DoublyNode
}

func NewDoublyLinkedList() *DoublyLinkedList {
	return new(DoublyLinkedList)
}

func (list *DoublyLinkedList) Add(data int) {
	newNode := &DoublyNode{data: data}
	if list.head == nil {
		list.head = newNode
	} else {
		current := list.head
		for current.next != nil {
			current = current.next
		}
		newNode.previous = current
		current.next = newNode
	}
}

func (list *DoublyLinkedList) Display() {
	current := list.head
	for current != nil {
		fmt.Printf("%d<->", current.data)
		current = current.next
	}
	fmt.Println("nil")
}

func (list *DoublyLinkedList) GetNode(index int) *DoublyNode {
	if index < 0 {
		return nil
	}
	current := list.head
	for i := 0; i < index && current != nil; i++ {
		current = current.next
	}

	return current
}
