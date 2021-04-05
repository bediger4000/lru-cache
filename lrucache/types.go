package lru

type CacheItem struct {
	key   LRUItem
	data  LRUItem
	hash  uint64
	chain *CacheItem
	next  *CacheItem
	prev  *CacheItem
}

type LRUItem interface {
	Hash() uint64
	Equals(otherKey interface{}) bool
}

func NewCacheStringItem(str string) *CacheItem {
	datum := &StringData{data: str}
	datum.Hash()
	return &CacheItem{key: datum, data: datum}
}

func NewCacheItem(key LRUItem, data LRUItem) *CacheItem {
	return &CacheItem{key: key, data: data}
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

func (c *Cache) Get(key LRUItem) interface{} {
	// move to front of list
	return nil
}

func (c *Cache) Set(key LRUItem, value interface{}) {
	var item *CacheItem
	switch value.(type) {
	case string:
		item = NewCacheItem(key, NewStringData(value.(string)))
	}
	if c.table.Insert(item) {
		c.current++
		// move to front of list
	}
	if c.current > c.n {
		// delete least recently used item
	}
}
