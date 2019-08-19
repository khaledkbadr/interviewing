package randomizedset

import (
	"math/rand"
)

// RandomizedSet data structure that supports Insert, Remove, getRandom operations in O(1)
type RandomizedSet struct {
	// A map where keys are inserted vals and vlaues are indexes in vals
	valIdx map[int]int
	// slice of vals inserted
	vals []int
}

// Constructor initialize your data structure here
func Constructor() RandomizedSet {
	return RandomizedSet{
		valIdx: make(map[int]int),
		vals:   make([]int, 0),
}

// Insert inserts a value to the set. Returns true if the set did not already contain the specified element
func (rs *RandomizedSet) Insert(val int) bool {
	if _, ok := rs.valIdx[val]; ok {
		return false
	}

	rs.vals = append(rs.vals, val)
	rs.valIdx[val] = len(rs.vals) - 1
	return true
}

// Remove removes a value from the set. returns true if the set contained the specified element
func (rs *RandomizedSet) Remove(val int) bool {
	idx, ok := rs.valIdx[val]
	if !ok {
		return false
	}

	// move the last value to deleted value index in vals
	lastVal := rs.vals[len(rs.vals)-1]
	rs.vals[idx] = lastVal
	rs.vals = rs.vals[:len(rs.vals)-1] // compact the slice
	rs.valIdx[lastVal] = idx           // update the last val index in valIdx map

	// delete value from rs.valIdx map
	delete(rs.valIdx, val)

	return true
}

// GetRandom Get a random element from the set.
func (rs *RandomizedSet) GetRandom() int {
	idx := rand.Intn(len(rs.vals))
	return rs.vals[idx]
}
