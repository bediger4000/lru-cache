package lru

type CacheItem struct {
	key   LRUItem
	data  LRUItem
	hash  uint32
	chain *CacheItem // hashtable chain
	next  *CacheItem // most-recently-used list
	prev  *CacheItem
}

// LRUItem interface allows a Cache data structure to hold
// data of any type that has keys of any type.
type LRUItem interface {
	Hash() uint32
	Equals(otherKey interface{}) bool
}

// StringData pointers and their methods fit interface LRUItem
type StringData struct {
	hash uint32
	data string
}

// Cache embodies the Least Recently Used cache asked
// for in the problem statement.
type Cache struct {
	table       *hashtable
	mostrecent  *CacheItem
	leastrecent *CacheItem
	n           int // max number of items in cache
	current     int // current number of items in cache
}
