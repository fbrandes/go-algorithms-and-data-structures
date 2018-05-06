package basic

type Node struct {
	item StackItem
	next Node
}

type Queue struct {
	size  int
	first Node
	last  Node
}

func (q *Queue) New() *Queue {
	q.size = 0
	return q
}

func (q *Queue) IsEmpty() bool {
	return q.size == 0
}

func (q *Queue) Size() int {
	return q.size
}

func (q *Queue) Peek() StackItem {
	// TODO check if Queue empty
	return q.first.item
}

func (q *Queue) enqueue(item StackItem) {

}