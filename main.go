package main

import (
	"fmt"
	lru "lru/lrucache"
	"os"
)

func main() {
	cacheSize := 5
	cache := lru.NewCache(cacheSize)
	unique := 0
	for _, str := range os.Args[1:] {
		datum := lru.NewStringData(str)
		fmt.Printf("New datum %q\n", str)
		if cache.Set(datum, str) {
			unique++
			if unique >= cacheSize {
				cache.PrintUse()
			}
		}
	}
	cache.PrintUse()
	fmt.Printf("Inserted %d unique items into cache\n", unique)
	for _, str := range os.Args[1:] {
		datum := lru.NewStringData(str)
		if d := cache.Get(datum); d != nil {
			fmt.Printf("Found %+v in cache\n", d)
		} else {
			fmt.Printf("did not find %q in cache\n", str)
		}
	}
}
