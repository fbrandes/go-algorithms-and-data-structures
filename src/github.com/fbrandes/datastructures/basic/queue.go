package basic

// based on: https://github.com/phf/go-queue/blob/master/queue/queue.go
// TODO finish implementation

type Queue struct {
	values []interface{}
	first  int
	last   int
	length int
}

func (q *Queue) New() *Queue {
	return new(Queue).Init()
}

func (q *Queue) Init() *Queue {
	q.values = make([]interface{}, 1)
	q.first, q.last, q.length = 0, 0, 0
	return q
}

func (q *Queue) Length() int {
	return q.length
}

func (q *Queue) IsEmpty() bool {
	return q.Length() == 0
}

func (q *Queue) IsFull() bool {
	return q.length == len(q.values)
}

func (q *Queue) isSparse() bool {
	return 1 < q.length && q.Length() < len(q.values)/4
}

func (q *Queue) Resize(size int) {
	resized := make([]interface{}, size)
	if q.first < q.last {
		copy(resized, q.values[q.first:q.last])
	} else {
		n := copy(resized, q.values[q.first:])
		copy(resized[n:], q.values[:q.last])
	}
	q.values = resized
	q.first = 0
	q.last = q.Length()
}

func (q *Queue) PeekFirst() interface{} {
	return q.values[q.first]
}

func (q *Queue) PeekLast() interface{} {
	return q.values[q.last]
}

func (q *Queue) Enqueue(v interface{}) {

}
