package minpriority

type Item struct {
	Value    string
	Priority int
}

type MinPriorityQueue []*Item

func NewMinPriority() *MinPriorityQueue {
	return new(MinPriorityQueue)
}

func (pq MinPriorityQueue) Len() int {
	return len(pq)
}

func (pq MinPriorityQueue) Less(i, j int) bool {
	return pq[i].Priority < pq[j].Priority
}

func (pq MinPriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *MinPriorityQueue) Push(x interface{}) {
	item := x.(*Item)
	*pq = append(*pq, item)
}

func (pq *MinPriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]

	return item
}
