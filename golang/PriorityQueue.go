package utils

import (
	"strconv"
	"sync"
)

type PriorityQueue struct {
	queues   [8]*Queue
	lock     *sync.Cond
	running  bool
	shutdown bool
	name     string
}

func NewPriorityQueue() *PriorityQueue {
	pq := &PriorityQueue{}
	pq.lock = sync.NewCond(&sync.Mutex{})
	pq.running = true

	for i := 0; i < 8; i++ {
		pq.queues[i] = newQueue()
	}
	return pq
}

func (pq *PriorityQueue) SetName(name string) {
	pq.name = name
}

func (pq *PriorityQueue) Push(any interface{}, priority int) error {
	if priority > 7 || priority < 0 {
		panic("Priority value range is 0..7 and cannot be:" + strconv.Itoa(priority))
	}
	pq.lock.L.Lock()
	if pq.shutdown || !pq.running {
		pq.lock.L.Unlock()
		return nil
	}
	err := pq.queues[priority].Push(any)
	pq.lock.Broadcast()
	pq.lock.L.Unlock()
	return err
}

func (pq *PriorityQueue) Pop() interface{} {
	for ; pq.running; {
		pq.lock.L.Lock()
		if pq.queues[7].size > 0 {
			elem := pq.queues[7].Pop()
			pq.lock.L.Unlock()
			return elem
		} else if pq.queues[6].size > 0 {
			elem := pq.queues[6].Pop()
			pq.lock.L.Unlock()
			return elem
		} else if pq.queues[5].size > 0 {
			elem := pq.queues[5].Pop()
			pq.lock.L.Unlock()
			return elem
		} else if pq.queues[4].size > 0 {
			elem := pq.queues[4].Pop()
			pq.lock.L.Unlock()
			return elem
		} else if pq.queues[3].size > 0 {
			elem := pq.queues[3].Pop()
			pq.lock.L.Unlock()
			return elem
		} else if pq.queues[2].size > 0 {
			elem := pq.queues[2].Pop()
			pq.lock.L.Unlock()
			return elem
		} else if pq.queues[1].size > 0 {
			elem := pq.queues[1].Pop()
			pq.lock.L.Unlock()
			return elem
		} else if pq.queues[0].size > 0 {
			elem := pq.queues[0].Pop()
			pq.lock.L.Unlock()
			return elem
		} else {
			pq.lock.Wait()
			pq.lock.L.Unlock()
		}
	}
	pq.lock.L.Lock()
	defer pq.lock.L.Unlock()
	pq.shutdown = true
	pq.lock.Broadcast()
	return nil
}

func (pq *PriorityQueue) Shutdown() {
	pq.lock.L.Lock()
	pq.running = false
	pq.lock.Broadcast()
	pq.lock.L.Unlock()
}
