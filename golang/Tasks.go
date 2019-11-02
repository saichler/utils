package utils

import (
	"sync"
)

type Job struct {
	maxParallelTasks int
	finishCount      int
	batchCount       int
	totalTasks       int
	queue            []JobTask
	l                *sync.Cond
	listener         JobListener
}

type JobTask interface {
	Run()
}

type JobListener interface {
	Finished(JobTask)
}

func NewJob(maxParallel int, listener JobListener) *Job {
	job := &Job{}
	job.queue = make([]JobTask, 0)
	job.maxParallelTasks = maxParallel
	job.l = sync.NewCond(&sync.Mutex{})
	job.listener = listener
	return job
}

func (job *Job) AddTask(task JobTask) {
	job.l.L.Lock()
	defer job.l.L.Unlock()
	job.queue = append(job.queue, task)
	job.totalTasks++
}

func (job *Job) Run() {
	for job.finishCount < job.totalTasks {
		toRun := make([]JobTask, 0)
		job.l.L.Lock()
		for i := 0; i < job.maxParallelTasks; i++ {
			task := job.queue[0]
			job.queue = job.queue[1:]
			toRun = append(toRun, task)
			if len(job.queue) == 0 {
				break
			}
		}
		job.batchCount = len(toRun)
		for _, task := range toRun {
			go job.runTask(task)
		}
		job.l.Wait()
		job.l.L.Unlock()
	}
}

func (job *Job) runTask(task JobTask) {
	task.Run()
	job.l.L.Lock()
	defer job.l.L.Unlock()
	if job.listener != nil {
		defer job.listener.Finished(task)
	}
	job.finishCount++
	job.batchCount--
	if job.batchCount == 0 {
		job.l.Broadcast()
	}
}
