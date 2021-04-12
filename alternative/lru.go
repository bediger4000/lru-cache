package main

import (
	"container/list"
	"fmt"
	"os"
)

type Cache struct {
	size    int
	current int
	table   map[string]*list.Element
	lru     *list.List
}

func main() {
	cache := NewCache(5)
	for _, str := range os.Args[1:] {
		cache.Set(str, str)
	}
	fmt.Printf("Cache contains %d items\n", cache.current)
	fmt.Println("LRU list:")
	for node := cache.lru.Front(); node != nil; node = node.Next() {
		data := node.Value.(string)
		fmt.Printf("%q\n", data)
	}
	fmt.Println("Table:")
	for key, item := range cache.table {
		fmt.Printf("Key %q, value %q\n", key, item.Value.(string))
	}
	for _, str := range os.Args[1:] {
		value := cache.Get(str)
		if value != "" {
			if value != str {
				fmt.Printf("Found %q as value for key %q\n", value, str)
				continue
			}
			fmt.Printf("Found %q\n", str)
		} else {
			fmt.Printf("Did not find %q\n", str)
		}
	}
}

func NewCache(n int) *Cache {
	return &Cache{
		size:  n,
		table: make(map[string]*list.Element),
		lru:   list.New(),
	}
}

func (c *Cache) Get(key string) string {
	if item, ok := c.table[key]; ok {
		c.lru.MoveToFront(item)
		return item.Value.(string)
	}
	return ""
}

func (c *Cache) Set(key string, value string) {
	var item *list.Element
	var ok bool
	if item, ok = c.table[key]; !ok {
		item = c.lru.PushFront(value)
		c.table[key] = item
		c.current++
	} else {
		c.lru.MoveToFront(item)
	}
	if c.current > c.size {
		item := c.lru.Back()
		c.lru.Remove(item) // this might be O(n) in length of list
		key := item.Value.(string)
		delete(c.table, key)
		c.current--
	}
}
