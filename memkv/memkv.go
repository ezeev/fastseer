package memkv

import (
	"sync"
	"time"
)

// idea from https://stackoverflow.com/questions/25484122/map-with-ttl-option-in-go

type item struct {
	value      interface{}
	lastAccess int64
}

type MemKV struct {
	m map[string]*item
	l sync.Mutex
}

func New(ln int, maxTTL int) (m *MemKV) {
	m = &MemKV{m: make(map[string]*item, ln)}
	go func() {
		for now := range time.Tick(time.Second) {
			m.l.Lock()
			for k, v := range m.m {
				if now.Unix()-v.lastAccess > int64(maxTTL) {
					delete(m.m, k)
				}
			}
			m.l.Unlock()
		}
	}()
	return
}

func (m *MemKV) Delete(key string) {
	m.l.Lock()
	delete(m.m, key)
	m.l.Unlock()
}

func (m *MemKV) Len() int {
	return len(m.m)
}

func (m *MemKV) Put(k string, v interface{}) {
	m.l.Lock()
	it, ok := m.m[k]
	if !ok {
		it = &item{value: v}
		m.m[k] = it
	}
	it.lastAccess = time.Now().Unix()
	m.l.Unlock()
}

func (m *MemKV) Get(k string) (v interface{}) {
	m.l.Lock()
	if it, ok := m.m[k]; ok {
		v = it.value
		it.lastAccess = time.Now().Unix()
	}
	m.l.Unlock()
	return

}
