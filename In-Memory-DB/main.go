package main

import (
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

func getValueByField(field string, items []Item) *string {
	for _, item := range items {
		if item.Field == field {
			return &item.Value
		}
	}

	return nil
}

func checkStringValue(str *string) string {
	if str != nil {
		return *str
	}

	return "None"
}

func main() {
	db := NewInMemoryDB()

	db.Set("A", "B", "C")
	db.Set("A", "D", "E")

	value := db.Get("A", "B")
	println(checkStringValue(value))

	value = db.Get("A", "D")
	println(checkStringValue(value))

	db.Delete("A", "B")

	value = db.Get("A", "B")
	println(checkStringValue(value))

	value = db.Get("A", "D")
	println(checkStringValue(value))
}
