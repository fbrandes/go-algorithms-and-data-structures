package advanced

import (
	"github.com/cheekybits/genny/generic"
	"sync"
)

// source: https://flaviocopes.com/golang-data-structure-dictionary/

type Key generic.Type
type Value generic.Type

type Dictionary struct {
	values map[Key]Value
	lock sync.RWMutex
}

func (d *Dictionary) Get(key Key) Value {
	d.lock.RLock()
	defer d.lock.RUnlock()
	return d.values[key]
}

func (d *Dictionary) Insert(key Key, value Value) {
	d.lock.Lock()
	defer d.lock.Unlock()
	if d.values == nil {
		d.values = make(map[Key]Value)
	}
	d.values[key] = value
}

func (d *Dictionary) Delete(key Key) bool {
	d.lock.Lock()
	defer d.lock.Unlock()
	_, ok := d.values[key]
	if ok {
		delete(d.values, key)
	}
	return ok
}

func (d *Dictionary) Contains(key Key) bool {
	d.lock.RLock()
	defer d.lock.RUnlock()
	_, ok := d.values[key]
	return ok
}

func (d *Dictionary) Size() int {
	d.lock.RLock()
	defer d.lock.RUnlock()
	return len(d.values)
}

func (d *Dictionary) Clear() {
	d.lock.Lock()
	defer d.lock.Unlock()
	d.values = make(map[Key]Value)
}

func (d *Dictionary) Keys() []Key {
	d.lock.RLock()
	defer d.lock.RUnlock()
	var keys []Key
	for i := range d.values {
		keys = append(keys, i)
	}
	return keys
}

func (d *Dictionary) Values() []Value {
	d.lock.RLock()
	defer d.lock.RUnlock()
	var values []Value
	for i := range d.values {
		values = append(values, d.values[i])
	}
	return values
}

