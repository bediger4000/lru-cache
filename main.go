package main

import (
	"fmt"
	lru "lru/lrucache"
	"os"
)

func main() {
	cache := lru.NewCache(20)
	for _, str := range os.Args[1:] {
		datum := lru.NewStringData(str)
		cache.Set(datum, str)
	}
	fmt.Printf("%+v\n", cache)
}
