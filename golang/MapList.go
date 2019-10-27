package utils

import "sync"

type MapList struct {
	m   map[interface{}]int
	l   []interface{}
	mtx *sync.Mutex
}

func NewMapList() *MapList {
	ml := &MapList{}
	ml.m = make(map[interface{}]int)
	ml.l = make([]interface{}, 0)
	ml.mtx = &sync.Mutex{}
	return ml
}

func (ml *MapList) Put(key, value interface{}) {
	ml.mtx.Lock()
	defer ml.mtx.Unlock()
	index := len(ml.l)
	ml.m[key] = index
	ml.l = append(ml.l, value)
}

func (ml *MapList) Get(key interface{}) interface{} {
	ml.mtx.Lock()
	defer ml.mtx.Unlock()
	index, ok := ml.m[key]
	if !ok {
		return nil
	}
	return ml.l[index]
}

func (ml *MapList) Del(key interface{}) interface{} {
	ml.mtx.Lock()
	defer ml.mtx.Unlock()
	index, ok := ml.m[key]
	if !ok {
		return nil
	}
	v := ml.l[index]
	ml.l[index] = nil
	return v
}

func (ml *MapList) List() []interface{} {
	ml.mtx.Lock()
	defer ml.mtx.Unlock()
	result := make([]interface{}, 0)
	for _, v := range ml.l {
		if v != nil {
			result = append(result, v)
		}
	}
	return result
}

func (ml *MapList) Map() map[interface{}]interface{} {
	ml.mtx.Lock()
	defer ml.mtx.Unlock()
	result := make(map[interface{}]interface{})
	for k, i := range ml.m {
		v := ml.l[i]
		if v != nil {
			result[k] = v
		}
	}
	return result
}
