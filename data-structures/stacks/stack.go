package stacks

import "errors"

type Stack struct {
	items []int
}

//Newstack creates an instance of Stack
func Newstack(length int) *Stack {
	return &Stack{items: make([]int, length)}
}

//Push adds an item onto the top of the stack
func (s *Stack) Push(item int) {
	s.items = append(s.items, item)
}

//Pop removes and returns the top item from the stack
func (s *Stack) Pop() (int, error) {
	if len(s.items) == 0 {
		return -1, errors.New("stack is empty")
	}
	lastIndex := len(s.items) - 1
	popped := s.items[lastIndex]
	s.items = s.items[:lastIndex]

	return popped, nil
}

//IsEmpty checks if the stack is empty
/*func (s *Stack) IsEmpty() bool {
	if len(s.items) == 0 {
		return true
	}
	return false
}*/

//Peek returns the top item from the stack
func (s *Stack) Peek() (int, error) {
	if len(s.items) == 0 {
		return -1, errors.New("stack is empty")
	}
	return s.items[len(s.items)-1], nil
}
