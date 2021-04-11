package lru

func NewCacheStringItem(str string) *CacheItem {
	datum := &StringData{data: str}
	return &CacheItem{key: datum, data: datum, hash: datum.Hash()}
}

func NewCacheItem(key LRUItem, data LRUItem) *CacheItem {
	return &CacheItem{key: key, data: data, hash: key.Hash()}
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
func (s *StringData) Hash() uint32 {
	if s.hash == 0 {
		var hash uint32 = 5381
		for _, b := range []byte(s.data) {
			hash = ((hash << 5) + hash) + uint32(b)
		}
		s.hash = hash
	}
	return s.hash
}
