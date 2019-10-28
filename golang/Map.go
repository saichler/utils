package utils

import "sync"

type Map struct {
	m   map[interface{}]interface{}
	mtx *sync.Mutex
}

func NewMap() *Map {
	ml := &Map{}
	ml.m = make(map[interface{}]interface{})
	ml.mtx = &sync.Mutex{}
	return ml
}

func (ml *Map) Put(key, value interface{}) {
	ml.mtx.Lock()
	defer ml.mtx.Unlock()
	ml.m[key] = value
}

func (ml *Map) Get(key interface{}) interface{} {
	ml.mtx.Lock()
	defer ml.mtx.Unlock()
	return ml.m[key]
}

func (ml *Map) Del(key interface{}) interface{} {
	ml.mtx.Lock()
	defer ml.mtx.Unlock()
	v := ml.m[key]
	delete(ml.m, key)
	return v
}

func (ml *Map) Map() map[interface{}]interface{} {
	ml.mtx.Lock()
	defer ml.mtx.Unlock()
	result := make(map[interface{}]interface{})
	for k, v := range ml.m {
		result[k] = v
	}
	return result
}
