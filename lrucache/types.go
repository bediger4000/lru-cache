package lru

import "fmt"

type CacheItem struct {
	key   LRUItem
	data  LRUItem
	hash  uint64
	chain *CacheItem // hashtable chain
	next  *CacheItem // most-recently-used list
	prev  *CacheItem
}

type LRUItem interface {
	Hash() uint64
	Equals(otherKey interface{}) bool
}

func NewCacheStringItem(str string) *CacheItem {
	datum := &StringData{data: str}
	return &CacheItem{key: datum, data: datum, hash: datum.Hash()}
}

func NewCacheItem(key LRUItem, data LRUItem) *CacheItem {
	return &CacheItem{key: key, data: data, hash: key.Hash()}
}

type StringData struct {
	hash uint64
	data string
}

func NewStringData(str string) *StringData {
	return &StringData{data: str}
}

func (s *StringData) Equals(other interface{}) bool {
	if s != nil {
		switch other.(type) {
		case *StringData:
			return s.data == other.(*StringData).data
		}
	}
	return false
}

func (s *StringData) Data() string {
	return s.data
}

// Hash method of StringData: DJB2 hash function
// extremely unlikely that the DJB2 hash of a string has value 0,
// but if it does, this recalculates zero every invocation.
func (s *StringData) Hash() uint64 {
	if s.hash == 0 {
		var hash uint64 = 5381
		for _, b := range []byte(s.data) {
			hash = ((hash << 5) + hash) + uint64(b)
		}
		s.hash = hash
	}
	return s.hash
}

type Cache struct {
	table       *hashtable
	mostrecent  *CacheItem
	leastrecent *CacheItem
	n           int // max number of items in cache
	current     int // current number of items in cache
}

func NewCache(n int) *Cache {
	return &Cache{
		table: NewTable(n),
		n:     n,
	}
}

func (c *Cache) Get(key LRUItem) interface{} {
	dataItem := c.table.Lookup(key)

	if dataItem == nil {
		return nil
	}

	if c.mostrecent == nil {
		c.mostrecent = dataItem
		c.leastrecent = dataItem
		return dataItem
	}

	c.updateMostRecent(dataItem)

	return dataItem.data
}

func (c *Cache) updateMostRecent(node *CacheItem) {

	fmt.Printf("Move data %q to most recently used\n",
		node.data.(*StringData).data,
	)

	if node == c.leastrecent {
		c.leastrecent = c.leastrecent.prev
	}

	if node == c.mostrecent {
		return
	}

	// chop it out of list
	if node.prev != nil {
		node.prev.next = node.next
	}
	if node.next != nil {
		node.next.prev = node.prev
	}

	// replace head
	node.next = c.mostrecent
	c.mostrecent.prev = node
	node.prev = nil
	c.mostrecent = node
}

func (c *Cache) Set(key LRUItem, value interface{}) bool {
	var item *CacheItem
	switch value.(type) {
	case string:
		item = NewCacheItem(key, NewStringData(value.(string)))
	}
	newinsert := false
	if node, inserted := c.table.Insert(item); inserted {
		c.current++
		newinsert = true
		// put on head of list
		item.next = c.mostrecent
		if c.mostrecent != nil {
			c.mostrecent.prev = item
			c.mostrecent = item
		} else {
			c.mostrecent = item
			c.leastrecent = item
		}
	} else {
		// a duplicate, node points to it
		c.updateMostRecent(node)
	}
	if c.current > c.n {
		// delete least recently used item
		tmp := c.leastrecent
		c.leastrecent = c.leastrecent.prev
		c.leastrecent.next = nil

		dkey := NewStringData(tmp.data.(*StringData).data)
		c.table.Delete(dkey)
		c.current--
	}
	return newinsert
}

func (c *Cache) PrintUse() {

	fmt.Printf("%d items in hashtable\n", c.table.Size())
	for node := c.mostrecent; node != nil; node = node.next {
		fmt.Printf("Data %q at %p: next %p, prev %p\n",
			node.data.(*StringData).data, node, node.next, node.prev,
		)
	}
	fmt.Printf("Least recent %q, next %p, prev %p\n",
		c.leastrecent.data.(*StringData).data,
		c.leastrecent.next,
		c.leastrecent.prev,
	)
}
