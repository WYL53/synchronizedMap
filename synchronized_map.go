package synchronizedMap

import (
	"sync"
)

type synchronizedMap struct {
	m       map[interface{}]interface{}
	rwMutex *sync.RWMutex
}

func New() *synchronizedMap {
	return &synchronizedMap{
		m:       make(map[interface{}]interface{}),
		rwMutex: &sync.RWMutex{},
	}
}

func (this *synchronizedMap) Get(key interface{}) interface{} {
	this.rwMutex.RLock()
	defer this.rwMutex.RUnlock()
	return this.m[key]
}

func (this *synchronizedMap) Set(key, value interface{}) {
	this.rwMutex.Lock()
	defer this.rwMutex.Unlock()
	this.m[key] = value
}

func (this *synchronizedMap) Each(f func(k, v interface{})) {
	this.rwMutex.RLock()
	mtemp := make(map[interface{}]interface{}, len(this.m))
	for key, value := range this.m {
		mtemp[key] = value
	}
	this.rwMutex.RUnlock()

	for key, value := range mtemp {
		f(key, value)
	}
}

func (this *synchronizedMap) Len() int {
	this.rwMutex.RLock()
	defer this.rwMutex.RUnlock()
	return len(this.m)
}

func (this *synchronizedMap) IsContain(key interface{}) bool {
	this.rwMutex.RLock()
	defer this.rwMutex.RUnlock()
	_, ok := this.m[key]
	return ok
}

func (this *synchronizedMap) Clear() {
	this.rwMutex.Lock()
	defer this.rwMutex.Unlock()
	this.m = make(map[interface{}]interface{})
}
