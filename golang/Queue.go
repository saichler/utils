package utils

type Queue struct {
	internalQueue []interface{}
	size          int
}

func newQueue() *Queue {
	q := &Queue{}
	q.internalQueue = make([]interface{}, 0)
	return q
}

func (q *Queue) Push(any interface{}) error {
	q.internalQueue = append(q.internalQueue, any)
	q.size++
	return nil
}

func (q *Queue) Pop() interface{} {
	if q.size > 0 {
		elem := q.internalQueue[0]
		q.internalQueue = q.internalQueue[1:]
		q.size--
		return elem
	}
	return nil
}
