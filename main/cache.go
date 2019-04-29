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
	return entry.value
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
			newTail := cache.tail.pre
			newTail.next = nil
			cache.tail = newTail
		} else {
			cache.size++
		}
	}
	cache.head = entryValue
}

type entry struct {
	key interface{}
	value interface{}
	pre *entry
	next *entry
}

func CreateNewCache(cap int) *LRUCache {
	cache := &LRUCache{cacheMap:make(map[interface{}]*entry), cap:cap}
	return cache
}


