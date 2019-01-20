package utils

import "sync"

type SyncQueue struct {
	internalQueue []interface{}
	size int
	cond *sync.Cond
}

func NewSyncQueue() *SyncQueue {
	q:=&SyncQueue{}
	q.internalQueue = make([]interface{},0)
	q.cond = sync.NewCond(&sync.Mutex{})
	return q
}

func (q *SyncQueue) Push(any interface{}) error {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()
	q.internalQueue = append(q.internalQueue,any)
	q.size++
	q.cond.Broadcast()
	return nil
}

func (q *SyncQueue) Pop() interface{} {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()
	return q.PopNoSync()
}

func (q *SyncQueue) PopNoSync() interface{} {
	if q.size==0 {
		q.cond.Wait()
	}
	if q.size>0 {
		elem := q.internalQueue[0]
		q.internalQueue = q.internalQueue[1:]
		q.size--
		return elem
	}
	return nil
}

func (q *SyncQueue) Size() int {
	return q.size
}

func (q *SyncQueue) Cond() *sync.Cond {
	return q.cond
}