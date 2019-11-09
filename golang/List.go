package utils

import "sync"

type List struct {
	internalList []interface{}
	lock         *sync.Cond
}

func NewList() *List {
	l := &List{}
	l.internalList = make([]interface{}, 0)
	l.lock = sync.NewCond(&sync.Mutex{})
	return l
}

func (l *List) Add(any interface{}) {
	l.lock.L.Lock()
	l.internalList = append(l.internalList, any)
	l.lock.L.Unlock()
	l.lock.Broadcast()
}

func (l *List) Get(pos int) interface{} {
	l.lock.L.Lock()
	defer l.lock.L.Unlock()
	if pos < 0 || len(l.internalList) <= pos {
		return nil
	}
	return l.internalList[pos]
}

func (l *List) Del(pos int) interface{} {
	l.lock.L.Lock()
	defer l.lock.L.Unlock()
	if pos < 0 || len(l.internalList) <= pos {
		return nil
	}
	elem := l.internalList[pos]
	aside := l.internalList[0:pos]
	zside := l.internalList[pos+1:]
	l.internalList = aside
	l.internalList = append(l.internalList, zside...)
	return elem
}

func (l *List) List(pos int) interface{} {
	l.lock.L.Lock()
	defer l.lock.L.Unlock()
	result := make([]interface{}, len(l.internalList))
	for i, v := range l.internalList {
		result[i] = v
	}
	return result
}

func (l *List) Size() int {
	l.lock.L.Lock()
	defer l.lock.L.Unlock()
	return len(l.internalList)
}
