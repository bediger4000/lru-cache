package main

import (
	"fmt"
	lru "lru/lrucache"
	"os"
)

func main() {
	cache := lru.NewCache(20)
	unique := 0
	for _, str := range os.Args[1:] {
		datum := lru.NewStringData(str)
		if cache.Set(datum, str) {
			unique++
		}
	}
	fmt.Printf("%+v\n", cache)
	fmt.Printf("Inserted %d unique items into cache\n", unique)
}
