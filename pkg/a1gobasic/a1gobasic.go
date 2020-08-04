package main

import (
	"fmt"
)

// Declare package level variable for storing map
var data map[string]string

func init() {
	data = make(map[string]string) // Initialise map with make
}

// Add given key and value to map
func addData(k, v string) {
	if _, ok := data[k]; !ok {
		fmt.Println("Given key not exists", k)
		fmt.Println("So adding the entry to map")
		data[k] = v
	} else {
		fmt.Println("Given key already exists", k)
	}
}

// Update given value for given key in map
func updateData(k, v string) {
	if _, ok := data[k]; ok {
		fmt.Println("Given key exists", k)
		fmt.Println("So updating the key value in the map")
		data[k] = v
	} else {
		fmt.Println("Given key not exists", k)
	}
}

// Get value of given key in a map
func getData(k string) string {
	if _, ok := data[k]; ok {
		fmt.Println("Given key exists", k)
		return data[k]
	} else {
		fmt.Println("Given key not exists", k)
	}
	return ""
}

// Get all values of map
func getAllData() ([]string, []string) {
	allValues := make([]string, len(data))
	allKeys := make([]string, len(data))
	for k, v := range data {
		allKeys = append(allKeys, k)
		allValues = append(allValues, v)
	}
	return allKeys, allValues
}

func deleteData(k string) {
	if _, ok := data[k]; ok {
		fmt.Println("Given key is exists", k)
		fmt.Println("So deleting the entry from map")
		delete(data, k)
	} else {
		fmt.Println("Given key not exists", k)
	}
}

func main() {

	// initialize the data with key/values
	data["key1"] = "value1"
	data["key2"] = "value2"
	data["key3"] = "value3"
	data["key4"] = "value4"

	// Test getData function - Positive
	key2value := getData("key2")
	fmt.Println("Positive::Get key2 value: ", key2value)
	// Test get function - Negative
	key5value := getData("key5")
	fmt.Println("Negative::Get key5 value: ", key5value)
	fmt.Println("")

	// Test addData function - Positive, add new key/value
	fmt.Println("Positive::Before add data: ", data)
	addData("key5", "value5")
	fmt.Println("Positive::After add data: ", data)
	// Test add function - Negative, add already exist key/value
	fmt.Println("Negative::Before add data: ", data)
	addData("key4", "value4")
	fmt.Println("Negative::After add data: ", data)
	fmt.Println("")

	// Test updateData function - Positive, update existing key/value
	fmt.Println("Positive::Before update data: ", data)
	updateData("key2", "value22")
	fmt.Println("Positive::After update data: ", data)
	// Test add function - Negative, update non-existing key/value
	fmt.Println("Negative::Before update data: ", data)
	updateData("key6", "value6")
	fmt.Println("Negative::After update data: ", data)
	fmt.Println("")

	// Test deleteData function - Positive, delete existing key/value
	fmt.Println("Positive::Before delete data: ", data)
	deleteData("key3")
	fmt.Println("Positive::After delete data: ", data)
	// Test add function - Negative, delete non-existing key/value
	fmt.Println("Negative::Before delete data: ", data)
	deleteData("key7")
	fmt.Println("Negative::After delete data: ", data)
	fmt.Println("")

	// Test getAllData function
	allDataKeys, allDataValues := getAllData()
	fmt.Println("All data keys: ", allDataKeys)
	fmt.Println("All data values: ", allDataValues)

}