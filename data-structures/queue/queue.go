package queue

import "errors"

type Queue struct {
	items []int
}

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) Enqueue(item int) {
	q.items = append(q.items, item)
}

func (q *Queue) Front() (int, error) {
	if len(q.items) == 0 {
		return -1, errors.New("the queue is empty")
	}
	front := q.items[0]
	return front, nil
}

func (q *Queue) Dequeue() (int, error) {
	if len(q.items) == 0 {
		return -1, errors.New("queue is empty")
	}
	dequeued := q.items[0]
	q.items = q.items[1:]

	return dequeued, nil
}
