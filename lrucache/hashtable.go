package lru

type hashtable struct {
	bucketcount int
	buckets     []*CacheItem
}

func NewTable(n int) *hashtable {
	count := n
	h := &hashtable{
		bucketcount: count,
		buckets:     make([]*CacheItem, count),
	}
	return h
}

// return true on insert, false on finding a duplicate
func (h *hashtable) Insert(item *CacheItem) bool {
	bucketIndex := int(item.hash) % h.bucketcount

	for node := h.buckets[bucketIndex]; node != nil; node = node.chain {
		if node.hash == item.hash {
			if node.data.Equals(item.data) {
				// found a duplicate
				return false
			}
		}
	}
	if node := h.buckets[bucketIndex]; node != nil {
		item.chain = node
		h.buckets[bucketIndex] = item
	} else {
		h.buckets[bucketIndex] = item
	}
	return true
}

// return true on delete, false when not finding key
func (h *hashtable) delete(key LRUItem) bool {
	return false
}
