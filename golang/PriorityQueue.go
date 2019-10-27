package utils

import (
	"strconv"
	"sync"
	"time"
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
	pq.lock.L.Unlock()
	pq.lock.Broadcast()
	return err
}

func (pq *PriorityQueue) Pop() interface{} {
	pq.lock.L.Lock()
	defer pq.lock.L.Unlock()

	for ; pq.running; {
		if pq.queues[7].size > 0 {
			return pq.queues[7].Pop()
		} else if pq.queues[6].size > 0 {
			return pq.queues[6].Pop()
		} else if pq.queues[5].size > 0 {
			return pq.queues[5].Pop()
		} else if pq.queues[4].size > 0 {
			return pq.queues[4].Pop()
		} else if pq.queues[3].size > 0 {
			return pq.queues[3].Pop()
		} else if pq.queues[2].size > 0 {
			return pq.queues[2].Pop()
		} else if pq.queues[1].size > 0 {
			return pq.queues[1].Pop()
		} else if pq.queues[0].size > 0 {
			return pq.queues[0].Pop()
		} else {
			pq.lock.Wait()
		}
	}
	pq.shutdown = true
	return nil
}

func (pq *PriorityQueue) Shutdown() {
	pq.running = false
	for i := 0; i < 100; i++ {
		pq.lock.Broadcast()
	}
	time.Sleep(time.Second / 5)
	if pq.shutdown {
		Info("Priority Queue " + pq.name + " was shutdown properly.")
	} else {
		Error("Priority Queue " + pq.name + " was not able to shutdown properly!")
		//panic("")
	}
}
