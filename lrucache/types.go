package lru

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

type StringData struct {
	hash uint64
	data string
}

type Cache struct {
	table       *hashtable
	mostrecent  *CacheItem
	leastrecent *CacheItem
	n           int // max number of items in cache
	current     int // current number of items in cache
}
