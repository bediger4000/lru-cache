package lru

func (p *CacheItem) Remove() {
	if p.prev != nil {
		p.prev.next = p.next
	}
	if p.next != nil {
		p.next.prev = p.prev
	}
}

func (p *CacheItem) InsertBefore(node *CacheItem) {
	if p != nil {
		node.prev = p.prev
	}
	node.next = p
	if node.next != nil {
		node.next.prev = node
	}
}
