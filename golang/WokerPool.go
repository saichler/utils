package utils

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

type WorkerPool struct {
	running    bool
	maxWorkers int
	working    int
	queue      []Task
	l          *sync.Cond
	lastReport int64
}

type Task interface {
	Run()
}

func NewWorkerPool(maxWorkers int) *WorkerPool {
	wp := &WorkerPool{}
	wp.queue = make([]Task, 0)
	wp.l = sync.NewCond(&sync.Mutex{})
	wp.maxWorkers = maxWorkers
	return wp
}

func (wp *WorkerPool) AddTask(task Task) {
	wp.l.L.Lock()
	defer wp.l.L.Unlock()
	wp.queue = append(wp.queue, task)
	wp.l.Broadcast()
}

func (wp *WorkerPool) Start() {
	wp.running = true
	go wp.executeTasks()
}

func (wp *WorkerPool) Stop() {
	wp.l.L.Lock()
	defer wp.l.L.Unlock()
	wp.running = false
	wp.l.Broadcast()
}

func (wp *WorkerPool) WaitForEmptyQueue() {
	wp.l.L.Lock()
	defer wp.l.L.Unlock()
	for len(wp.queue) > 0 {
		wp.l.Wait()
	}
}

func (wp *WorkerPool) executeTasks() {
	for wp.running {
		wp.l.L.Lock()
		if time.Now().Unix()-wp.lastReport >= 15 {
			wp.lastReport = time.Now().Unix()
			if len(wp.queue) > 0 {
				fmt.Println("Queue Size=" + strconv.Itoa(len(wp.queue)))
			}
		}
		if len(wp.queue) == 0 || wp.working == wp.maxWorkers {
			wp.l.Wait()
		}
		var task Task
		if len(wp.queue) > 0 && wp.working < wp.maxWorkers {
			task = wp.queue[0]
			wp.queue = wp.queue[1:]
			wp.working++
		}
		wp.l.L.Unlock()
		if task != nil {
			task.Run()
			wp.l.L.Lock()
			wp.working--
			wp.l.Broadcast()
			wp.l.L.Unlock()
		}
	}
}
