package lru

func (p *CacheItem) Remove() {
	if p.prev != nil {
		p.prev.next = p.next
	}
	if p.next != nil {
		p.next.prev = p.prev
	}
}

// insert node before p in the linked list
// does not return node, since to call it you need
// both p and node. p probably points to head of linked list.
func (p *CacheItem) InsertBefore(node *CacheItem) {
	if p != nil {
		node.prev = p.prev
	}
	node.next = p
	if node.next != nil {
		node.next.prev = node
	}
}
