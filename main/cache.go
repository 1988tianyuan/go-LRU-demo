package main


type LRUCache struct {
	head *entry
	tail *entry
	cap  int
	size int
	cacheMap map[interface{}]*entry
}

func (cache *LRUCache) GetValue(key interface{}) interface{} {
	entry := cache.cacheMap[key]
	if entry == nil {
		return nil
	} else {
		previousHead := cache.head
		if entry == previousHead {
			return entry.value
		}
		cache.lruHandling(entry, previousHead)
		cache.head = entry
	}
	return entry.value
}

func (cache *LRUCache) lruHandling(entryValue *entry, previousHead *entry)  {
	previousPre := entryValue.pre
	previousNext := entryValue.next
	entryValue.next = previousHead
	previousHead.pre = entryValue
	if previousNext == nil {	// entryValue is the previous tail
		previousPre.next = nil
		cache.tail = previousPre
	} else {	// entryValue is not the previous tail
		previousPre.next = previousNext
		previousNext.pre = previousPre
	}
	entryValue.pre = nil
}

func (cache *LRUCache) SaveValue(key interface{}, value interface{}) {
	cacheMap := cache.cacheMap
	previousHead := cache.head
	var entryValue *entry
	entryValue = cacheMap[key]
	if entryValue != nil {
		entryValue = cacheMap[key]
		entryValue.value = value
		if entryValue == previousHead {
			return
		}
		cache.lruHandling(entryValue, previousHead)
	} else {
		entryValue = &entry{key: key, value:value}
		cacheMap[key] = entryValue
		if cache.size == 0 {
			cache.tail = entryValue
		} else {
			entryValue.next = previousHead
			previousHead.pre = entryValue
		}
		if cache.size == cache.cap {
			cache.Remove(cache.tail.key)
			//newTail := cache.tail.pre
			//newTail.next = nil
			//cache.tail = newTail
		} else {
			cache.size++
		}
	}
	cache.head = entryValue
}

func (cache *LRUCache) Remove(key interface{})  {
	entry := cache.cacheMap[key]
	if entry == nil {
		return
	}
	previousPre := entry.pre
	previousNext := entry.next
	if entry == cache.head {
		cache.head = previousNext
		if previousNext != nil {
			previousNext.pre = nil
		}
	}
	if entry == cache.tail {
		cache.tail = previousPre
		if previousPre != nil {
			previousPre.next = nil
		}
	}
}

type entry struct {
	key interface{}
	value interface{}
	pre *entry
	next *entry
}

func CreateNewCache(cap int) *LRUCache {
	cache := &LRUCache{cacheMap: make(map[interface{}]*entry), cap: cap}
	return cache
}

func (it *Iterator) Next() (key interface{}, value interface{}) {
	next := it.next
	it.next = next.next
	return next.key, next.value
}

func (it *Iterator) HasNext() bool {
	return it.next != nil
}

func (cache *LRUCache) GetIterator() *Iterator {
	it := &Iterator{cache.head}
	return it
}

type Iterator struct {
	next *entry
}