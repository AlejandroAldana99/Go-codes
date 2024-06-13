package main

func checkStringValue(str *string) string {
	if str != nil {
		return *str
	}

	return "None"
}

func main() {
	// db := NewInMemoryDB()

	// db.Set("A", "B", "C")
	// db.Set("A", "D", "E")

	// value := db.Get("A", "B")
	// println(checkStringValue(value))

	// value = db.Get("A", "D")
	// println(checkStringValue(value))

	// db.Delete("A", "B")

	// value = db.Get("A", "B")
	// println(checkStringValue(value))

	// value = db.Get("A", "D")
	// println(checkStringValue(value))
}
