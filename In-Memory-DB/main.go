package main

import (
	"sync"
)

type Item struct {
	Field string
	Value string
}

type InMemoryImpl struct {
	mu        sync.Mutex
	ItemsList map[string][]Item
}

func NewInMemoryDB() *InMemoryImpl {
	return &InMemoryImpl{
		ItemsList: make(map[string][]Item),
	}
}

func (db *InMemoryImpl) Set(key, field, value string) {
	db.mu.Lock()
	defer db.mu.Unlock()

	if items, exists := db.ItemsList[key]; exists {
		db.ItemsList[key] = append(items, Item{Field: field, Value: value})
	} else {
		db.ItemsList[key] = []Item{{Field: field, Value: value}}
	}
}

func (db *InMemoryImpl) Get(key, field string) *string {
	db.mu.Lock()
	defer db.mu.Unlock()

	if items, exists := db.ItemsList[key]; exists {
		return getValueByField(field, items)
	}

	return nil
}

func (db *InMemoryImpl) Delete(key, field string) {
	db.mu.Lock()
	defer db.mu.Unlock()

	if items, exists := db.ItemsList[key]; exists {
		for i, item := range items {
			if item.Field == field {
				db.ItemsList[key] = append(items[:i], items[i+1:]...)
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

func checkStringValue(str *string) string {
	if str != nil {
		return *str
	}

	return "None"
}
