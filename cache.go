package bezier

import (
	"fmt"
)

type Cache struct {
	Cap  int
	Size int
	head *node
	data []*node
}

type node struct {
	Value KV
	Key   string
	next  *node
}

type KV interface {
	Key() string
}

func NewCache(capacity int) *Cache {
	ca := &Cache{
		Cap:  capacity,
		Size: -1,
		head: nil,
		data: make([]*node, capacity),
	}
	for i := 0; i < capacity; i++ {
		ca.data[i] = &node{}
	}
	return ca
}

func (c *Cache) Put(i KV) {
	c.Size++
	if c.head == nil {
		c.head = c.data[0]
		c.head.Value = i
		c.head.Key = i.Key()
		c.head.next = nil
		return
	}
	loc := (c.Size) % c.Cap
	cur := c.head
	newHead := c.data[loc]
	newHead.Value = i
	newHead.Key = i.Key()
	newHead.next = cur
	c.head = newHead
}

func (c *Cache) Get(key string) KV {
	cur := c.head
	start := cur
	for cur != nil {
		if cur.Key == key {
			return cur.Value
		}
		cur = cur.next
		if start == cur {
			break
		}
	}
	return nil
}

func (c *Cache) String() string {
	ret := ""
	size := c.Size + 1
	if size >= c.Cap {
		size = c.Cap
	}
	ret += fmt.Sprintf("cap:%d, size:%d\n", c.Cap, size)
	cur := c.head
	start := cur
	for cur != nil {
		ret += fmt.Sprintf("%+v -> ", cur.Key)
		cur = cur.next
		if start == cur {
			break
		}
	}
	ret += "\n"
	return ret
}
