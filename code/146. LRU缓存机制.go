package code

type LRUCache struct {
	Capacity	int
	Count		int
	Length		int
	Size		int
	Caches		[]*Cache
}

type Cache struct {
	Key			int
	Value		int
	Count		int
	Next		*Cache
}


func Constructor(capacity int) LRUCache {
	lc := LRUCache{
		Capacity:	capacity,
		Count:		0,
		Length:		0,
		Caches:		make([]*Cache, capacity / 10 + 1),
		Size:		capacity / 10 + 1,
	}
	return lc
}


func (this *LRUCache) Get(key int) int {
	cache := this.Caches[key % this.Size]
	result := -1
	var temp *Cache
	for cache != nil {
		if cache.Key == key {
			this.Count++
			cache.Count = this.Count
			result = cache.Value
			if cache.Next != nil {
				t := cache
				if temp == nil {
					this.Caches[key % this.Size] = cache.Next
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


func (this *LRUCache) Put(key int, value int)  {
	cache := this.Caches[key % this.Size]
	result := -1
	var temp *Cache
	for cache != nil {
		if cache.Key == key {
			this.Count++
			cache.Count = this.Count
			cache.Value = value
			result = value
			if cache.Next != nil {
				t := cache
				if temp == nil {
					this.Caches[key % this.Size] = cache.Next
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
	if this.Length == this.Capacity {
		min := this.Count + 1
		pin := 0
		for i := 0; i < this.Size; i++ {
			if this.Caches[i] != nil && this.Caches[i].Count < min {
				min = this.Caches[i].Count
				pin = i
			}
		}
		if temp != nil && this.Caches[pin].Key == temp.Key {
			temp = nil
		}
		this.Caches[pin] = this.Caches[pin].Next
	} else {
		this.Length++
	}
	this.Count++
	t := &Cache {
		Key:	key,
		Value:	value,
		Count:	this.Count,
	}
	if temp == nil {
		this.Caches[key % this.Size] = t
	} else {
		temp.Next = t
	}
}
