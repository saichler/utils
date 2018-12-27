package utils

import "sync"

type Queue struct {
	inbox []interface{}
	lock *sync.Cond
}

func NewQueue() *Queue {
	q:=&Queue{}
	q.inbox = make([]interface{},0)
	q.lock = sync.NewCond(&sync.Mutex{})
	return q
}

func (q *Queue) Push(any interface{}) error {
	q.lock.L.Lock()
	defer q.lock.L.Unlock()
	q.inbox = append(q.inbox,any)
	q.lock.Broadcast()
	return nil
}

func (q *Queue) Pop() interface{} {
	q.lock.L.Lock()
	defer q.lock.L.Unlock()

	var elem interface{}
	for ;elem==nil; {
		if len(q.inbox)==0 {
			q.lock.Wait()
		} else {
			elem = q.inbox[0]
			q.inbox=q.inbox[1:]
		}
	}

	return elem
}