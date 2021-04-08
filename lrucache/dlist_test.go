package lru

import "testing"

func Test_Remove(t *testing.T) {
	dlist := &CacheItem{}
	dlist.Remove()
}
func Test_Remove2(t *testing.T) {
	head := &CacheItem{data: NewStringData("A")}
	node := &CacheItem{data: NewStringData("B")}
	head.next = node
	node.prev = head
	node2 := &CacheItem{data: NewStringData("C")}
	node.next = node2
	node2.prev = node
	node2.Remove()
	if node.next != nil {
		t.Fatal("failed to remove node from list")
	}
}
func Test_Move1(t *testing.T) {
	head := &CacheItem{data: NewStringData("A")}
	node := &CacheItem{data: NewStringData("B")}
	head.next = node
	node.prev = head
	node2 := &CacheItem{data: NewStringData("C")}
	node.next = node2
	node2.prev = node
	node2.Remove()
	if node.next != nil {
		t.Fatal("failed to remove node from list")
	}
	head.InsertBefore(node2)

	if node2.next != head {
		t.Fatal("failed to move node to head of list")
	}
}
func Test_Move2(t *testing.T) {
	head := &CacheItem{data: NewStringData("A")}
	node := &CacheItem{data: NewStringData("B")}
	head.next = node
	node.prev = head
	node2 := &CacheItem{data: NewStringData("C")}
	node.next = node2
	node2.prev = node

	node.Remove()
	if head.next != node2 || node2.prev != head {
		t.Fatal("failed to remove node from list")
	}

	head.InsertBefore(node)

	if node.next != head {
		t.Fatal("failed to move node to head of list")
	}
}
