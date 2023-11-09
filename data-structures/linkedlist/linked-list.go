package linkedlist

import "fmt"

type SinglyNode struct {
	data int
	next *SinglyNode
}

type SinglyLinkedList struct {
	head *SinglyNode
}

func NewSinglyLinkedList() *SinglyLinkedList {
	return new(SinglyLinkedList)
}

func (list *SinglyLinkedList) Add(data int) {
	newNode := &SinglyNode{data: data}
	if list.head == nil {
		list.head = newNode
	} else {
		current := list.head
		for current.next != nil {
			current = current.next
		}
		current.next = newNode
	}
}

func (list *SinglyLinkedList) Display() {
	current := list.head
	for current != nil {
		fmt.Printf("%d->", current.data)
		current = current.next
	}
	fmt.Println(nil)
}

func (list *SinglyLinkedList) GetNode(index int) *SinglyNode {
	if index < 0 {
		return nil
	}
	current := list.head
	for i := 0; i < index && current != nil; i++ {
		current = current.next
	}
	return current
}
