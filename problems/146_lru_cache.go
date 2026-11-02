package problems

//Time complexity of operations ~ O(1), memory - O(n)

type Node struct {
	val  int
	key  int
	prev *Node
	next *Node
}

type LRUCache struct {
	hashMap map[int]*Node
	cap     int
	first   *Node
	last    *Node
}

func ConstructorLRUCache(capacity int) LRUCache {
	return LRUCache{
		hashMap: make(map[int]*Node),
		cap:     capacity,
		first:   nil,
		last:    nil,
	}
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.hashMap[key]; ok {
		if node != this.last {
			if node == this.first {
				this.first = node.next
			}
			prev := node.prev
			next := node.next
			if prev != nil {
				prev.next = next
			}
			if next != nil {
				next.prev = prev
			}
			node.prev = this.last
			node.next = nil
			this.last.next = node
			this.last = node
		}
		return node.val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.hashMap[key]; ok {
		node.val = value
		if node != this.last {
			if node == this.first {
				this.first = node.next
			}
			prev := node.prev
			next := node.next
			if prev != nil {
				prev.next = next
			}
			if next != nil {
				next.prev = prev
			}
			node.prev = this.last
			node.next = nil
			this.last.next = node
			this.last = node
		}
	} else {
		newNode := Node{
			val:  value,
			key:  key,
			next: nil,
			prev: nil,
		}
		if this.last != nil {
			this.last.next = &newNode
			newNode.prev = this.last
		}
		this.last = &newNode
		if this.first == nil {
			this.first = &newNode
		}
		this.hashMap[key] = &newNode
		if len(this.hashMap) > this.cap {
			delete(this.hashMap, this.first.key)
			this.first = this.first.next
			this.first.prev = nil
		}
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := ConstructorLRUCache(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
