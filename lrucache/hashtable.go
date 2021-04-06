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

// return LRUItem container on finding it, nil otherwise
func (h *hashtable) Lookup(key LRUItem) *CacheItem {
	bucketIndex := int(key.Hash()) % h.bucketcount
	keyHash := key.Hash()
	for node := h.buckets[bucketIndex]; node != nil; node = node.chain {
		if node.hash == keyHash {
			if node.data.Equals(key) {
				return node
			}
		}
	}
	return nil
}

// return true on insert, false on finding a duplicate
// On duplicate, also return the node it found, to be able
// to move it to the head of the list
func (h *hashtable) Insert(item *CacheItem) (*CacheItem, bool) {
	bucketIndex := int(item.hash) % h.bucketcount

	for node := h.buckets[bucketIndex]; node != nil; node = node.chain {
		if node.hash == item.hash {
			if node.data.Equals(item.data) {
				// found a duplicate
				return node, false
			}
		}
	}
	if node := h.buckets[bucketIndex]; node != nil {
		item.chain = node
		h.buckets[bucketIndex] = item
	} else {
		h.buckets[bucketIndex] = item
	}
	return h.buckets[bucketIndex], true
}

// return true on delete, false when not finding key
func (h *hashtable) delete(key LRUItem) bool {
	return false
}
