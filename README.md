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

I ended up with 460 lines of Go,
implementing a single-chained hashtable (not a Go map)
and a doubly-linked list,
as well as the cache's `get` and `set` methods.
The cache uses an interface, so multiple types of keys and data
could exist, but I only implemented a string key and data type.
I wrote by data types rather than using standard library or package
data types so that I could ensure O(1) operation.

Creating the LRU cache sets up the buckets of the hashtable,
sets n.
The number of buckets in the hashtable should be about 1/10 of n.
That would mean that a full cache (n items in it),
if the hashing function is good,
the chains of items would average a length of 10.
That's said to be an optimal length.

The doubly-linked list lets the cache keep track of the "least recently used" property.
When a "get" operation finds an item in the hashtable,
the code moves that item from somewhere in the doubly-linked list
to the front of that list: it's the most recently used.

I chose to do "invasive" data structures:
the `CacheItem` struct has elements that allow code to
put a `CacheItem` instance into a single-chained hash table
and a doubly-linked list at the same time.
This is contrary to Object Oriented dogma,
where the data would be referred to by a hash table data structure,
and a doubly-linked list node structure.
The programmer would use standard libraries of "Queues"
and "Dictionaries" to hold the data.

I chose to hold the data in structures that match a Go interface
I named `LRUItem`. The phrasing of the problem statement
didm't give any clues about the types of data or keys:
either data or key could be a string, or an integer,
or a floating point number.
I only implemented a string data type,
using it for both cache key and data.
Because the key and data can be any type,
having the `CacheItem` type be a node in a singly-linked
hashtable chain, and a doubly-linked "most recently used" list
makes the most sense.
The cache doesn't have to have `n` hashtable list nodes,
and another `n` LRU doubly-linked list nodes.
The programmer doesn't have to keep track of 2 container structs
per data item in the cache,
and the algorithm doesn't have to deal with potential lack of locality
for 2 container data structs that refer to the same cached data.
There's probably also some memory savings:
there's no Go type-header for a single single-chain-pointer,
and another Go type-header for a doubly-linke list pointer.

The speed of lookups in the single-chained hash table
is dependent to a large extent on the hashing function
used to distribute data items over the hash table buckets.
I used the well-known [DJB2](http://www.cse.yorku.ca/~oz/hash.html)
hashing function hoping that items get distributed over the
number of buckets (item chains), and that there are very
few duplicates.
DJB2 hashing appears to work well.

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

A single-chaining hash table is easy to code,
but may not constitute the best hash table for a particular data item
or key type.

Defining the "least-recently-used" property by order in a fixed-size
doubly-length list might not be as fast as one would wish,
or as easy to code.
Using a doubly-linked list as a circular buffer might work.
The "tail" of the list is just `head.next`, so once the cache
is full, a new datum replaces the datum at `head.next`,
and becomes the new head. 
It's possible that a fixed-size array or slice,
treated as a circular buffer would work well.
My rationalizations about a since `CacheItem` container could
tip in favor of a Go slice of pointers, even for very large
numbers of cache items.
Coding might be easier, and memory usage might go down.

## Around the web

* [Standard library-based](https://anothercasualcoder.blogspot.com/2018/11/least-recently-used-lru-cache-by-google.html)
implementation. Not at all sure what language this is in.
* [C++ STL](https://www.geeksforgeeks.org/lru-cache-implementation/)
implementation.
* [Python](https://codereview.stackexchange.com/questions/225788/least-recently-used-cache-daily-coding-practice)

All 3 of these other implementations use the same hashtable and queue cache implementation.
They all use some standard library data types to do the hashtable and doubly-linked list.
They all have less than 20% of the lines of code I ended up with.
I just might have failed this interview question.

To see if Go can do an equally succint version of an LRU cache
if the programmer exploits standard library code,
I wrote a 78-line [version](alternative/lru.go) that exploits
Go's "map" data type, and a standard package doubly-linked list
container.
Even here, I used the doubly-linked list container `*list.Element`
as the data type stored in the map,
mainly due to the `list.Remove` and `list.MoveToFront` methods
of the doubly-linked list package.

This version does have drawbacks:

* Not clear how efficient the map data type is.
* Not clear that the `list.Remove` item is O(1).
It might actually walk the list to remove the item,
although that probably isn't true.

These objections might hold for any of the LRU caches that
exploit standard library or package or template code,
and boil down to objecting to a black box implementation.
