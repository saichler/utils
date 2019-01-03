package utils

import "sync"

type List struct {
	internalList []interface{}
	lock *sync.Cond
}

func NewList() *List {
	l:=&List{}
	l.internalList = make([]interface{},0)
	l.lock = sync.NewCond(&sync.Mutex{})
	return l
}

func (l *List) Add(any interface{}) {
	l.lock.L.Lock()
	l.internalList = append(l.internalList,any)
	l.lock.L.Unlock()
	l.lock.Broadcast()
}

func (l *List) Get(pos int) interface{} {
	l.lock.L.Lock()
	defer l.lock.L.Unlock()
	return l.internalList[pos]
}

func (l *List) Size() int {
	l.lock.L.Lock()
	defer l.lock.L.Unlock()
	return len(l.internalList)
}