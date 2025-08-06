package main

import (
	"sync"
	"time"
)

type KV struct {
	mu   sync.RWMutex
	data map[string][]byte
	ttl  map[string]int64
}

func NewKV() *KV {
	return &KV{
		data: map[string][]byte{},
		ttl:  make(map[string]int64),
	}
}

func (kv *KV) Set(key, val []byte) error {
	kv.mu.Lock()
	defer kv.mu.Unlock()
	kv.data[string(key)] = []byte(val)
	return nil
}

func (kv *KV) SetTTL(key []byte, seconds int64) {
	kv.mu.Lock()
	defer kv.mu.Unlock()
	skey := string(key)
	kv.ttl[skey] = time.Now().Unix() + seconds
}

func (kv *KV) Get(key []byte) ([]byte, bool) {
	kv.mu.Lock()
	defer kv.mu.Unlock()
	skey := string(key)

	if expireAt, ok := kv.ttl[skey]; ok {
		if time.Now().Unix() > expireAt {
			delete(kv.data, skey)
			delete(kv.ttl, skey)
			return nil, false
		}
	}

	val, ok := kv.data[skey]
	return val, ok
}

func (kv *KV) Delete(key []byte) int {
	kv.mu.Lock()
	defer kv.mu.Unlock()
	_, existed := kv.data[string(key)]
	if existed {
		delete(kv.data, string(key))
		return 1
	}
	return 0
}
