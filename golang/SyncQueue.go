package utils

import "sync"

type SyncQueue struct {
	internalQueue []interface{}
	cond          *sync.Cond
}

func NewSyncQueue() *SyncQueue {
	q := &SyncQueue{}
	q.internalQueue = make([]interface{}, 0)
	q.cond = sync.NewCond(&sync.Mutex{})
	return q
}

func (q *SyncQueue) Push(any interface{}) error {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()
	q.internalQueue = append(q.internalQueue, any)
	q.cond.Broadcast()
	return nil
}

func (q *SyncQueue) Pop() interface{} {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()
	for len(q.internalQueue) == 0 {
		q.cond.Wait()
	}
	elem := q.internalQueue[0]
	q.internalQueue = q.internalQueue[1:]
	return elem
}

func (q *SyncQueue) Size() int {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()
	return len(q.internalQueue)
}
