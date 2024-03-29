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

Famously, hash tables are O(1), at least amortized over many look-ups.
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

A single-chain hash table would be filled with data containers
that look like this:

```go
type CacheItem struct {
	key    LRUItem
	data   LRUItem
    hash   uint64
    chain *CacheItem
    next  *CacheItem
    prev  *CacheItem
}

type LRUItem interface {
	Hash() uint32
	Equals(otherKey LRUItem) bool
}
```

`CacheItem.chain` element points to the next item in the hash bucket,
while `CacheItem.prev` and `CacheItem.next` point to doubly-linked list items
used to determine use status.
The containers have elements that allow them to exist in 2
different data structures,
a linked list (for the single-chained hash table)
and a doubly-linked list, for the recent use status.

Defining an interface (I'm doing this in Go) for the keys
lets me ignore the key type until I need a key of some
particular type, `int` or `string` or whatever.
The algorithm for searching a hash chain can compare hash values
until it finds a numerically-equal hash value,
then call the `LRUItem.Equals` method to do an actual comparison.

I used the same named type `LRUItem` for both key and data.
I'm not sure this is a good idea, but for the time being, it works.

Is the LRU cache used by multiple threads?
Nothing in the problem statement about that,
so I'm assuming no.
Single-threaded use only.

### Algorithm Design

The LRU cache has:

1. Single-chaining hash table
2. Doubly-linked list
3. An int, n, the max number of items in the cache
4. An int representing the current number of items in the cache.

I ended up with 460 lines of Go,
implementing a single-chained hash table (not a Go map)
and a doubly-linked list,
as well as the cache's `get` and `set` methods.
The cache uses an interface, so multiple types of keys and data
could exist, but I only implemented a string key and data type.
I wrote by data types rather than using standard library or package
data types so that I could ensure O(1) operation.

Creating the LRU cache sets up the buckets of the hash table,
sets n.
The number of buckets in the hash table should be about 1/10 of n.
That would mean that a full cache (n items in it),
if the hashing function is good,
the chains of items would average a length of 10.
That's said to be an optimal length.

The doubly-linked list lets the cache keep track of the "least recently used" property.
When a "get" operation finds an item in the hash table,
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
didn't give any clues about the types of data or keys:
either data or key could be a string, or an integer,
or a floating point number.
I only implemented a string data type,
using it for both cache key and data.
Because the key and data can be any type,
having the `CacheItem` type be a node in a singly-linked
hash table chain, and a doubly-linked "most recently used" list
makes the most sense.
The cache doesn't have to have `n` hash table list nodes,
and another `n` LRU doubly-linked list nodes.
The programmer doesn't have to keep track of 2 container structs
per data item in the cache,
and the algorithm doesn't have to deal with potential lack of locality
for 2 container data structs that refer to the same cached data.
There's probably also some memory savings:
there's no Go type-header for a single single-chain-pointer,
and another Go type-header for a doubly-linked list pointer.

The speed of lookouts in the single-chained hash table
is dependent to a large extent on the hashing function
used to distribute data items over the hash table buckets.
I used the well-known [DJB2](http://www.cse.yorku.ca/~oz/hash.html)
hashing function hoping that items get distributed over the
number of buckets (item chains), and that there are very
few duplicates.
DJB2 hashing appears to work well,
but I did need to use Go's `uint32` type for hash values.
Apparently DJB2's good distribution depends on periodically overflowing
a 32-bit value.

#### set(key, value)

1. Create a new `CacheItem`
   * get hash from `key.Hash()`
   * set `CacheItem.data`
2. Increment the current number of items in the cache
   * if it's &gt; n, find least recently used container from doubly-linked list.
   * Remove least recently used container from doubly-linked list and hash table.
   * decrement the current number of items in the cache, it will have value of n
3. Add `CacheItem` container to hash table
4. Add `CacheItem` container to head (most recently used) end of doubly-linked list.

It should be possible to instantiate only n `CacheItem` containers.
If the cache is full, remove the least-recently-used item
from the doubly-linked list, and delete it from the hash table.
Reset key and hash value and data,
put on the head of the doubly-linked list,
and re-insert into the hash table under the new key.
I did not do this optimization.

#### get(key)

1. Get hash from `key.Hash()`
2. Find a container matching the hash in hash table.
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
My rationalizations about a single `CacheItem` container could
tip in favor of a Go slice of pointers, even for very large
numbers of cache items.
Coding might be easier, and memory usage might go down.

## Other Implementations

This problem appears in the "Hash Tables" chapter
of the [Daily Coding Problem book](https://www.amazon.com/Daily-Coding-Problem-exceptionally-interviews/dp/1793296634/ref=sr_1_3?dchild=1&keywords=daily+coding+problem&qid=1627421725&sr=8-3)
They show an approximately 65-line Python solution,
but only the class definitions, no invocation scaffolding.

* [Standard library-based](https://anothercasualcoder.blogspot.com/2018/11/least-recently-used-lru-cache-by-google.html)
implementation. Not at all sure what language this is in.
* [C++ STL](https://www.geeksforgeeks.org/lru-cache-implementation/)
implementation.
* [Python](https://codereview.stackexchange.com/questions/225788/least-recently-used-cache-daily-coding-practice)

All 3 of these other implementations use the same hash table and queue
cache implementation.
They all use some standard library data types to
do the hash table and doubly-linked list.
They all have less than 20% of
the lines of code I ended up with.  I just might have failed this
interview question.

To see if Go can do an equally succinct version of an LRU cache
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
