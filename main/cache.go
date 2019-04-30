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
		cache.size++
		if cache.size > cache.cap {
			cache.removeEntry(cache.tail)
		}
	}
	cache.head = entryValue
}

func (cache *LRUCache) removeEntry(entry *entry) interface{} {
	if entry == nil {
		return nil
	}
	value := entry.value
	delete(cache.cacheMap, entry.key)
	previousPre := entry.pre
	previousNext := entry.next
	entry.next = nil
	entry.pre = nil
	entry.value = nil
	cache.size--
	if entry == cache.head {
		cache.head = previousNext
		if previousNext != nil {
			previousNext.pre = nil
		}
		return value
	}
	if entry == cache.tail {
		cache.tail = previousPre
		if previousPre != nil {
			previousPre.next = nil
		}
		return value
	}
	previousPre.next = previousNext
	previousNext.pre = previousPre
	return value
}

func (cache *LRUCache) Remove(key interface{}) interface{} {
	entry := cache.cacheMap[key]
	return cache.removeEntry(entry)
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