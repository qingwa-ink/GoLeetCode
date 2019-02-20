package code

// LRUCache : 缓存数据存储结构
type LRUCache struct {
	Capacity int
	Count    int
	Length   int
	Size     int
	Caches   []*Cache
}

// Cache ： 缓存数据
type Cache struct {
	Key   int
	Value int
	Count int
	Next  *Cache
}

// Constructor : 初始化一个LRUCache对象
func Constructor(capacity int) LRUCache {
	lc := LRUCache{
		Capacity: capacity,
		Count:    0,
		Length:   0,
		Caches:   make([]*Cache, capacity/10+1),
		Size:     capacity/10 + 1,
	}
	return lc
}

// Get : 取出缓存数据
func (cat *LRUCache) Get(key int) int {
	cache := cat.Caches[key%cat.Size]
	result := -1
	var temp *Cache
	for cache != nil {
		if cache.Key == key {
			cat.Count++
			cache.Count = cat.Count
			result = cache.Value
			if cache.Next != nil {
				t := cache
				if temp == nil {
					cat.Caches[key%cat.Size] = cache.Next
				} else {
					temp.Next = cache.Next
				}
				for cache.Next != nil {
					cache = cache.Next
				}
				cache.Next = t
				t.Next = nil
			}
			break
		}
		temp = cache
		cache = cache.Next
	}
	return result
}

// Put : 存入数据
func (cat *LRUCache) Put(key int, value int) {
	cache := cat.Caches[key%cat.Size]
	result := -1
	var temp *Cache
	for cache != nil {
		if cache.Key == key {
			cat.Count++
			cache.Count = cat.Count
			cache.Value = value
			result = value
			if cache.Next != nil {
				t := cache
				if temp == nil {
					cat.Caches[key%cat.Size] = cache.Next
				} else {
					temp.Next = cache.Next
				}
				for cache.Next != nil {
					cache = cache.Next
				}
				cache.Next = t
				t.Next = nil
			}
			break
		}
		temp = cache
		cache = cache.Next
	}
	if result != -1 {
		return
	}
	for temp != nil && temp.Next != nil {
		temp = temp.Next
	}
	if cat.Length == cat.Capacity {
		min := cat.Count + 1
		pin := 0
		for i := 0; i < cat.Size; i++ {
			if cat.Caches[i] != nil && cat.Caches[i].Count < min {
				min = cat.Caches[i].Count
				pin = i
			}
		}
		if temp != nil && cat.Caches[pin].Key == temp.Key {
			temp = nil
		}
		cat.Caches[pin] = cat.Caches[pin].Next
	} else {
		cat.Length++
	}
	cat.Count++
	t := &Cache{
		Key:   key,
		Value: value,
		Count: cat.Count,
	}
	if temp == nil {
		cat.Caches[key%cat.Size] = t
	} else {
		temp.Next = t
	}
}
