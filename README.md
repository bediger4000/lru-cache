# Daily Coding Problem: Problem #848 [Hard]

This problem was asked by Google.

Implement an LRU (Least Recently Used) cache.
It should be able to be initialized with a cache size n,
and contain the following methods:

* `set(key, value)`: sets key to value.
If there are already n items in the cache and we are adding a new item,
then it should also remove the least recently used item.
* `get(key)`: gets the value at key.
If no such key exists, return null.

Each operation should run in O(1) time.

## Analysis

This is a design problem,
including both data structure and algorithm elements.

Famously, hashtables are O(1), at least amortized over many lookups.
The problem statement seems like a big hint.

The semantics of the `set()` method imply that a
fixed-size circular buffer
should be used to keep track of the time-of-use or order-of-use
of the n items in the cache.
Every time `get()` gets called, move that item to the head of
the circular buffer.
What about when the program calls `set()`?
Does that put the item at the head or tail of the LRU status buffer?

### Data Design

A single-chain hashtable would be filled with data containers
that look like this:

```go
type CacheItem struct {
	key    LRUKey
	data   interface{}
    hash   uint64
    chain *CacheItem
    next  *CacheItem
    prev  *CacheItem
}

type LRUKey interface {
	Hash() uint64
	Equals(otherKey LRUKey) bool
}
```

`CacheItem.chain` element points to the next item in the hashbucket,
while `CacheItem.prev` and `CacheItem.next` point to doubly-linked list items
used to determine use status.
The containers have elements that allow them to exist in 2
different data structures,
a linked list (for the single-chained hash table)
and a doubly-linked list, for the recent use status.

Defining an interface (I'm doing this in Go) for the keys
lets me ignore the key type until I need a key of some
particular type, `int` or `string` or whatever.

Is the LRU cache used by multiple threads?
Nothing in the problem statement about that,
so I'm assuming no.
Single-threaded use only.

### Algorithm Design

The LRU cache has:

1. Single-chaining hashtable
2. Doubly-linked list
3. An int, n, the max number of items in the cache
4. An int representing the current number of items in the cache.

Creating the LRU cache sets up the buckets of the hashtable,
sets n.

#### set(key, value)

1. Create a new `CacheItem`
   * get hash from `key.Hash()`
   * set `CacheItem.data`
2. Increment the current number of items in the cache
   * if it's &gt; n, find least recently used container from doubly-linked list.
   * Remove least recently used container from doubly-linked list and hashtable.
   * decrement the current number of items in the cache, it will have value of n
3. Add `CacheItem` container to hashtable
4. Add `CacheItem` container to head (most recently used) end of doubly-linked list.

#### get(key)

1. Get hash from `key.Hash()`
2. Find a container matching the hash in hashtable.
   * If it exists, move the container to the head of the doubly-linked list.
   * If it exists, return the `CacheItem.data`

## Interview Analysis

Unlike most Daily Programming Problems,
this one really does deserve a "[hard]" rating.
Several data structures, operations on each structure
when doing LRU cache operations,
probably a choice of data structure in a few places.
