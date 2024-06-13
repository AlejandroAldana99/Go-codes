package memory

import (
	"fmt"
	"strings"
	"sync"
)

type Item struct {
	Field     string
	Value     string
	Timestamp int
	Ttl       int
}

type InMemoryDBImpl struct {
	mu       sync.Mutex
	ItemList map[string][]Item
}

func NewInMemoryDB() *InMemoryDBImpl {
	return &InMemoryDBImpl{
		ItemList: make(map[string][]Item),
	}
}

func (db *InMemoryDBImpl) Set(key, field, value string) {
	db.mu.Lock()
	defer db.mu.Unlock()

	if items, exists := db.ItemList[key]; exists {
		db.ItemList[key] = append(items, Item{Field: field, Value: value})
	} else {
		db.ItemList[key] = []Item{{Field: field, Value: value}}
	}
}

func (db *InMemoryDBImpl) Get(key, field string) *string {
	db.mu.Lock()
	defer db.mu.Unlock()

	if items, exists := db.ItemList[key]; exists {
		return getValueByField(field, items)
	}

	return nil
}

func (db *InMemoryDBImpl) Delete(key, field string) {
	db.mu.Lock()
	defer db.mu.Unlock()

	if items, exists := db.ItemList[key]; exists {
		for i, item := range items {
			if item.Field == field {
				db.ItemList[key] = append(items[:i], items[i+1:]...)
				break
			}
		}
	}
}

func (db *InMemoryDBImpl) SetAt(key string, field string, value string, timestamp int) {
	db.mu.Lock()
	defer db.mu.Unlock()

	items := db.ItemList[key]
	for i, item := range items {
		if item.Field == field {
			items[i].Value = value
			items[i].Timestamp = timestamp
			return
		}
	}
	db.ItemList[key] = append(items, Item{Field: field, Value: value, Timestamp: timestamp})
}

func (db *InMemoryDBImpl) SetAtWithTtl(key string, field string, value string, timestamp int, ttl int) {
	db.mu.Lock()
	defer db.mu.Unlock()

	items := db.ItemList[key]
	for i, item := range items {
		if item.Field == field {
			items[i].Value = value
			items[i].Timestamp = timestamp
			items[i].Ttl = ttl
			return
		}
	}
	db.ItemList[key] = append(items, Item{Field: field, Value: value, Timestamp: timestamp, Ttl: ttl})
}

func (db *InMemoryDBImpl) GetAt(key string, field string, timestamp int) *string {
	db.mu.Lock()
	defer db.mu.Unlock()

	items := db.ItemList[key]
	for _, item := range items {
		if item.Field == field && (item.Ttl == 0 || timestamp <= item.Timestamp+item.Ttl) {
			return &item.Value
		}
	}
	return nil
}

func (db *InMemoryDBImpl) DeleteAt(key string, field string, timestamp int) bool {
	db.mu.Lock()
	defer db.mu.Unlock()

	items := db.ItemList[key]
	for i, item := range items {
		if item.Field == field {
			db.ItemList[key] = append(items[:i], items[i+1:]...)
			return true
		}
	}
	return false
}

func (db *InMemoryDBImpl) ScanAt(key string, timestamp int) []string {
	db.mu.Lock()
	defer db.mu.Unlock()

	var results []string
	items := db.ItemList[key]
	for _, item := range items {
		if item.Ttl == 0 || timestamp <= item.Timestamp+item.Ttl {
			results = append(results, fmt.Sprintf("%s(%s)", item.Field, item.Value))
		}
	}
	return results
}

func (db *InMemoryDBImpl) ScanByPrefixAt(key string, prefix string, timestamp int) []string {
	db.mu.Lock()
	defer db.mu.Unlock()

	var results []string
	items := db.ItemList[key]
	for _, item := range items {
		if strings.HasPrefix(item.Field, prefix) && (item.Ttl == 0 || timestamp <= item.Timestamp+item.Ttl) {
			results = append(results, fmt.Sprintf("%s(%s)", item.Field, item.Value))
		}
	}
	return results
}

func getValueByField(field string, items []Item) *string {
	for _, item := range items {
		if item.Field == field {
			return &item.Value
		}
	}

	return nil
}
