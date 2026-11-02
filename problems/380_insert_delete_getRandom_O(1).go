package problems

//Time complexity of operations ~ O(1), memory - O(n)

import "math/rand/v2"

type RandomizedSet struct {
	slice   []int
	hashMap map[int]int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		slice:   make([]int, 0),
		hashMap: make(map[int]int),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.hashMap[val]; ok {
		return false
	}
	this.slice = append(this.slice, val)
	this.hashMap[val] = len(this.slice) - 1
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	if index, ok := this.hashMap[val]; ok {
		lastNum := this.slice[len(this.slice)-1]
		this.slice[index] = lastNum
		this.hashMap[lastNum] = index
		this.slice = this.slice[:len(this.slice)-1]
		delete(this.hashMap, val)
		return true
	}
	return false
}

func (this *RandomizedSet) GetRandom() int {
	rand := rand.N(len(this.slice))
	return this.slice[rand]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
