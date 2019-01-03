package utils

import "sync"

type Queue struct {
	internalQueue []interface{}
	lock          *sync.Cond
}

func NewQueue() *Queue {
	q:=&Queue{}
	q.internalQueue = make([]interface{},0)
	q.lock = sync.NewCond(&sync.Mutex{})
	return q
}

func (q *Queue) Push(any interface{}) error {
	q.lock.L.Lock()
	q.internalQueue = append(q.internalQueue,any)
	q.lock.L.Unlock()
	q.lock.Broadcast()
	return nil
}

func (q *Queue) Pop() interface{} {
	q.lock.L.Lock()
	defer q.lock.L.Unlock()

	var elem interface{}
	for ;elem==nil; {
		if len(q.internalQueue)==0 {
			q.lock.Wait()
		} else {
			elem = q.internalQueue[0]
			q.internalQueue =q.internalQueue[1:]
		}
	}

	return elem
}