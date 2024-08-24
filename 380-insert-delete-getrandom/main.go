package main

import "math/rand"

type RandomizedSet struct {
	set  map[int]int
	list []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		set:  make(map[int]int),
		list: []int{},
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.set[val]; !ok {
		index := len(this.list)
		this.list = append(this.list, val)
		this.set[val] = index

		return true
	}
	return false
}

func (this *RandomizedSet) Remove(val int) bool {
	if index, ok := this.set[val]; ok {
		lastIndex := len(this.list) - 1
		lastValue := this.list[lastIndex]

		this.list[index] = lastValue
		this.set[lastValue] = index

		this.list = this.list[:lastIndex]
		delete(this.set, val)

		return true
	}
	return false
}

func (this *RandomizedSet) GetRandom() int {
	return this.list[rand.Intn(len(this.list))]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
