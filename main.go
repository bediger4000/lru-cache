package main

import (
	"fmt"
	lru "lru/lrucache"
)

func main() {
	cache := lru.NewCache(20)
	fmt.Printf("%+v\n", cache)
}
