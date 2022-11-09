package main

import "fmt"

const CACHE_SIZE = 5

type Hash map[string]*Node

// Structure corresponding to a single value
type Node struct {
	Val  string
	Next *Node
	Prev *Node
}

// Contains all nodes in a linked-list like structure
type Queue struct {
	Head   *Node
	Tail   *Node
	Length int
}

// Contains the queue of nodes and a hash map that facilitates searching for values
type Cache struct {
	Queue Queue
	Hash  Hash
}

// Initializes a new cache
func NewCache() Cache {
	return Cache{Queue: NewQueue(), Hash: Hash{}}
}

// Initializes a new queue
func NewQueue() Queue {
	head := &Node{}
	tail := &Node{}

	head.Next = tail
	tail.Prev = head

	return Queue{Head: head, Tail: tail}
}

func main() {
	fmt.Println("START CACHE")
	cache := NewCache()

	items := []string{"genesis", "exodus", "matthew", "john", "mark", "john", "revelation"}

	// add words to cache
	for _, word := range items {
		cache.Check(word)
		cache.Display()
	}
}

// Check if a value exists in cache and delete if so
// finally add to cache
func (c *Cache) Check(str string) {
	node := &Node{}

	if val, ok := c.Hash[str]; ok {
		node = c.Remove(val)
	} else {
		node = &Node{Val: str}
	}

	c.Add(node)
	c.Hash[str] = node
}

func (c *Cache) Display() {
	c.Queue.Display()
}

// Remove node by linking its left and right pointers
func (c *Cache) Remove(node *Node) *Node {
	fmt.Printf("Removing node: %s\n", node.Val)

	left := node.Prev
	right := node.Next

	left.Next = right
	right.Prev = left

	c.Queue.Length--

	delete(c.Hash, node.Val)

	return node
}

// Add node to head of queue (node should be at the right of the head)
func (c *Cache) Add(node *Node) {
	fmt.Printf("Adding node: %s\n", node.Val)

	temp := c.Queue.Head.Next

	c.Queue.Head.Next = node

	node.Prev = c.Queue.Head
	node.Next = temp

	temp.Prev = node

	c.Queue.Length++

	// This implements the least recently used (lru) feature
	// The lru node is at the left of the tail of queue so if cache is full, delete this node
	if c.Queue.Length > CACHE_SIZE {
		c.Remove(c.Queue.Tail.Prev)
	}
}

// Print queue values to stdout
func (q *Queue) Display() {
	node := q.Head.Next
	fmt.Printf("%d - [", q.Length)

	for i := 0; i < q.Length; i++ {
		fmt.Printf("{%s}", node.Val)

		if i < q.Length-1 {
			fmt.Printf("<-->")
		}
		node = node.Next
	}

	fmt.Println("]")
}
