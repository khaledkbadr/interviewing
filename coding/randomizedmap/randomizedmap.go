package randomizedmap

import (
	"math/rand"
)

// RandomizedMap Map that supports Insert, Remove, getRandom operations in O(1)
type RandomizedMap struct {
	// A map where keys are inserted keys and vlaues are indexes in vals
	kv map[int][]int
	// slice of vals inserted
	keys []int
}

// Constructor initialize your data structure here
func Constructor() RandomizedMap {
	return RandomizedMap{
		kv:   make(map[int][]int),
		keys: make([]int, 0),
	}
}

// Insert inserts a key, value to the map.
func (rm *RandomizedMap) Insert(key int, val int) {
	_, ok := rm.kv[key]
	// new
	if !ok {
		rm.keys = append(rm.keys, key)
		rm.kv[key] = []int{val, len(rm.keys) - 1}
	} else {
		rm.kv[key][0] = val
	}
}

// Remove removes a value from the map. returns true if the map contained the specified key
func (rm *RandomizedMap) Remove(key int) bool {
	v, ok := rm.kv[key]
	if !ok {
		return false
	}

	idx := v[1]

	// move the last value to deleted value index in keys
	lastKey := rm.keys[len(rm.keys)-1]
	rm.keys[idx] = lastKey
	rm.keys = rm.keys[:len(rm.keys)-1] // compact the slice
	rm.kv[lastKey][1] = idx            // update the last val index in kv map

	// delete value from rm.kv map
	delete(rm.kv, key)

	return true
}

// GetRandom Get a random value from the map.
func (rm *RandomizedMap) GetRandom() int {
	idx := rand.Intn(len(rm.keys))
	return rm.kv[rm.keys[idx]][0]
}
