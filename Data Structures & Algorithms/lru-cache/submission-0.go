type Node struct {
	Key int
	Val int
	Next *Node
	Prev *Node
}

type LRUCache struct {
    Capacity int
	Data map[int]*Node
	Head *Node
	Tail *Node
}

func Constructor(capacity int) LRUCache {
    head, tail := &Node{}, &Node{}
	head.Next = tail
	tail.Prev = head
	data := make(map[int]*Node, capacity)

	return LRUCache{
		Capacity: capacity,
		Data: data,
		Head: head,
		Tail: tail,
	}
}

func (this *LRUCache) AddToHead(node *Node) {
	node.Next = this.Head.Next //head node --> a --> tail
	node.Prev = this.Head
	this.Head.Next.Prev = node
	this.Head.Next = node
}

func (this *LRUCache) Remove(node *Node) {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.Data[key]
	if !ok {
		return -1
	}

	this.Remove(node)
	this.AddToHead(node)

	return node.Val
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.Data[key]; ok {
		this.Remove(node)
		this.AddToHead(node)

		node.Val = value

		return
	}

	if len(this.Data) >= this.Capacity {
		node := this.Tail.Prev
		this.Remove(node)
		
		delete(this.Data, node.Key)
	}

	new_node := &Node{
		Key: key,
		Val: value,
	}

	this.AddToHead(new_node)
	this.Data[key] = new_node
}
