package main

type MyHashSet struct {
	innerMap map[int]bool
}

func Constructor() MyHashSet {
	return MyHashSet{
		innerMap: make(map[int]bool),
	}
}

func (this *MyHashSet) Add(key int) {
	this.innerMap[key] = true
}

func (this *MyHashSet) Remove(key int) {
	delete(this.innerMap, key)
}

func (this *MyHashSet) Contains(key int) bool {
	return this.innerMap[key]
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
