package lru

type CacheItem struct {
	key   LRUKey
	data  LRUData
	hash  uint64
	chain *CacheItem
	next  *CacheItem
	prev  *CacheItem
}

type LRUKey interface {
	Hash() uint64
	Equals(otherKey LRUKey) bool
}

type LRUData interface {
	Equals(other LRUData) bool
}

type IntKey uint64

func (p IntKey) Hash() uint64 {
	return uint64(p)
}

type StringData struct {
	hash uint64
	data string
}

func (s *StringData) Equals(other *StringData) bool {
	if s != nil && other != nil {
		if s.data == other.data {
			return true
		}
	}
	return false
}

func (p IntKey) Equals(otherKey LRUKey) bool {
	switch otherKey.(type) {
	case IntKey:
		return p == otherKey.(IntKey)
	default:
		return uint64(p) == otherKey.Hash()
	}
	return false
}

type Cache struct {
	table   *hashtable
	head    *CacheItem
	tail    *CacheItem
	n       int // max number of items in cache
	current int // current number of items in cache
}

func NewCache(n int) *Cache {
	return &Cache{
		table: NewTable(n),
		n:     n,
	}
}

func (c *Cache) Get(key LRUKey) interface{} {
	// move to front of list
	return nil
}

func (c *Cache) Set(key LRUKey, value interface{}) {
	c.current++
}
