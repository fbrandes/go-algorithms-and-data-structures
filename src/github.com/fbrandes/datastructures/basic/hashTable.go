package basic

import (
	"github.com/cheekybits/genny/generic"
	"sync"
	"fmt"
)

type Key generic.Type

type Value generic.Type

type HashTable struct {
	items map[int]Value
	lock  sync.RWMutex
}

func hash(k Key) int {
	key := fmt.Sprintf("&s", k)
	hash := 0
	for i := 0; i < len(key); i++ {
		hash = 31*hash + int(key[i])
	}
	return hash
}

func (h *HashTable) Put(k Key, v Value) {
	h.lock.Lock()
	defer h.lock.Unlock()
	index := hash(k)

	if h.items == nil {
		h.items = make(map[int]Value)
	}
	h.items[index] = v
}

func (h *HashTable) Remove(k Key) {
	h.lock.Lock()
	defer h.lock.Unlock()
	index := hash(k)
	delete(h.items, index)
}

func (h *HashTable) Get(k Key) Value {
	h.lock.RLock()
	defer h.lock.RUnlock()
	index := hash(k)
	return h.items[index]
}

func (h *HashTable) Size() int {
	h.lock.RLock()
	defer h.lock.RUnlock()
	return len(h.items)
}
