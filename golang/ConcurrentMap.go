package utils

import "sync"

type ConcurrentMap struct {
	innerMap map[interface{}]interface{}
	mtx *sync.Mutex
}

func NewConcurrentMap() *ConcurrentMap {
	cm:=&ConcurrentMap{}
	cm.innerMap=make(map[interface{}]interface{})
	cm.mtx = &sync.Mutex{}
	return cm
}

func (cm *ConcurrentMap)Put(key,value interface{}) {
	cm.mtx.Lock()
	defer cm.mtx.Unlock()
	cm.innerMap[key]=value
}

func (cm *ConcurrentMap)Get(key interface{}) (interface{},bool) {
	cm.mtx.Lock()
	defer cm.mtx.Unlock()
	value,ok:= cm.innerMap[key]
	return value,ok
}

func (cm *ConcurrentMap)Del(key interface{}) bool {
	cm.mtx.Lock()
	defer cm.mtx.Unlock()
	_,ok:=cm.innerMap[key]
	delete(cm.innerMap,key)
	return ok
}

func (cm *ConcurrentMap)GetMap() map[interface{}]interface{} {
	cm.mtx.Lock()
	defer cm.mtx.Unlock()
	result:=make(map[interface{}]interface{})
	for k,v:=range cm.innerMap {
		result[k]=v
	}
	return result
}

