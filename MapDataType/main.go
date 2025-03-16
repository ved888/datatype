package main

import (
	"encoding/json"
	"fmt"
	"sort"
)

func main() {
	MapDeclarations()
	MakeNilMap()
	MergeMap()
	AssignAnother()
	x := CompareTwoMaps()
	fmt.Println("return compare two map========", x)
	MapStoreAnotherMapAsAValue()
	MapKeyToSlice()
	MapValuesToSlice()
	SortTheKeysOfAMap()
	SortTheValuesOfAMap()
	MapWhereTheKeyIsAStruct()
	MapWhereTheValueIsAStruct()
	ConvertAMapToAJSONstring()
	ConvertAJSONstringToAMap()
	FindTheKeyWithTheHighestValue()
	FindTheKeyWithTheLowestValue()
	RemoveMultipleKeysFromAMap()
	CountOccurrencesOfElementsUsingAMap()
	MapAsACache()
	MapWhereTheValueIsASlice()
	IteratingOverAMapOfSlices()
}

func MapDeclarations() {
	// Using map literal with shorthand
	x := map[string]string{
		"name": "ved prakash",
	}
	// Using make function with shorthand
	y := make(map[string]string)
	y["Name"] = "Rahul verma"
	// Using var keyword (must be initialized before use)
	var z map[string]string
	z = make(map[string]string) //initialized here
	z["name"] = "Ravi"
	z["age"] = "24"

	// Using var with make function directly
	var a = make(map[string]string)
	a["name"] = "Raman"
	// update the value of a key in a map
	a["name"] = "shyam"
	// delete a key-value pair from a map
	delete(a, "name")
	// Using var with map literal
	var b = map[string]string{
		"name": "Raghav",
	}
	// map contain duplicate keys using as valuse is slice
	c := map[string][]int{
		"mango":  {2, 5},
		"banana": {6, 9, 8},
	}

	for k, v := range c {
		fmt.Println("Key:", k, "Value:", v)
	}
	// Printing maps
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)
	fmt.Println("return only age", z["age"])
	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println("c:", c)
}

func MakeNilMap() {
	// methode
	myMap := map[string]int{
		"apple": 4,
		"mango": 7,
	}
	fmt.Println("Before Clear:", myMap)
	// Clear the map
	myMap = nil

	fmt.Println("After Clear:", myMap) // Output: map[]

	// methode 2
	x := make(map[string]string)
	x["apple"] = "fruit"
	x["valvo"] = "car"
	fmt.Println("Before Clear:", x)
	// Clear the map
	myMap = make(map[string]int) // Reinitialize the map

	fmt.Println("After Clear:", myMap) // Output: map[]
}

func MergeMap() {
	x := map[string]int{
		"apple": 4,
		"valvo": 9,
	}
	y := map[string]int{
		"computer": 4,
		"car":      7,
	}

	// Merging map2 into map1
	for key, valuse := range y {
		x[key] = valuse
	}
	fmt.Println("Merged Map:", x)

}

// 23.What happens when you assign one map variable to another?

func AssignAnother() {
	original := map[string]int{
		"apple":  9,
		"greaps": 100,
	}

	// Assigning one map to another
	copyMap := original
	// Modifying the new map
	copyMap["car"] = 10
	fmt.Println("Original Map:", original)
	fmt.Println("Copy Map:", copyMap)
}

func CompareTwoMaps() bool {

	myMap1 := map[string]string{
		"name": "ved",
		"age":  "28",
		"add":  "varanasi",
	}

	MyMap2 := map[string]string{
		"name":    "Rashul",
		"email":   "rahul@gmail.com",
		"contact": "09876123456",
	}

	// First, check if both maps have the same length
	if len(myMap1) != len(MyMap2) {
		return false
	}

	for key, valu := range myMap1 {
		value, exists := MyMap2[key]
		if !exists || valu != value {
			return false
		}
	}
	return true
}

func MapStoreAnotherMapAsAValue() {
	// Define a map where values are also maps
	var myMap = map[string]map[string]string{
		"emp1": {
			"name":  "ved",
			"email": "ved@gmail.com",
			"phone": "234567898",
		},
		"address": {
			"add1": "kardhana",
			"pin":  "8909",
			"dist": "varanasi",
		},
	}
	fmt.Println("map store anthore map as a values data", myMap)
	fmt.Println("print only address:=", myMap["address"]["add1"])

	myMap["address2"] = map[string]string{
		"add2":    "bodhapur",
		"state":   "uttar pradesh",
		"country": "India",
	}

	// Modify a value
	myMap["address2"]["add2"] = "lucknow"

	fmt.Println("print address2 data:=", myMap["address2"]["country"])
	fmt.Println("print map:=", myMap)
}

func MapKeyToSlice() {
	myMap := map[string]string{
		"name":  "ved",
		"add":   "varanasi",
		"phone": "12344456777",
	}

	keys := make([]string, 0, len(myMap))

	for key, _ := range myMap {
		keys = append(keys, key)
	}
	fmt.Println("map to slice", keys)
}

func MapValuesToSlice() {
	myMap := map[string]string{
		"name": "ved prakash",
		"Add":  "karadhana",
		"Cont": "98687984456",
	}
	val := make([]string, 0, len(myMap))
	for _, value := range myMap {
		val = append(val, value)
	}
	fmt.Println("map to slice as a value", val)
}

func SortTheKeysOfAMap() {
	myMap := map[string]string{
		"name": "ved",
		"add":  "varanasi",
		"cont": "9909090909",
	}
	// Step 1: Extract keys into a slice
	mapToslice := make([]string, 0, len(myMap))
	for key, _ := range myMap {
		mapToslice = append(mapToslice, key)
	}
	// Step 2: Sort the slice of keys
	sort.Strings(mapToslice)

	// Step 3: Iterate over the sorted keys
	fmt.Println("Sorted map by keys:")

	for _, key := range mapToslice {
		fmt.Printf("%s: %v\n", key, myMap[key])
	}
}

func SortTheValuesOfAMap() {
	myMap := map[string]string{
		"name": "ved prakash",
		"add":  "bodhapur",
		"cont": "98493759843",
	}
	// Convert map to a slice of key-value pairs
	type kv struct {
		Key   string
		Value string
	}
	var kvSlice []kv

	for key, value := range myMap {
		kvSlice = append(kvSlice, kv{key, value})
	}
	// Sort by values (ascending)
	sort.Slice(kvSlice, func(i, j int) bool {
		return kvSlice[i].Value < kvSlice[j].Value
	})

	// Print sorted key-value pairs
	fmt.Println("Sorted map by values:")
	for _, kv := range kvSlice {
		fmt.Printf("%s: %s\n", kv.Key, kv.Value)
	}
}

func MapWhereTheKeyIsAStruct() {
	type myStruct struct {
		Name string
		Add  string
	}

	myMap := map[myStruct]string{
		{Name: "ved", Add: "vns"}:       "software engineer",
		{Name: "Rahul", Add: "lucknow"}: "data scientist",
	}
	// add new values
	myMap[myStruct{Name: "Raj", Add: "vns"}] = "school boy"
	fmt.Println("map key as a struct==", myMap)
	//modified the values
	// make key like this way
	key := myStruct{Name: "Raj", Add: "vns"}
	myMap[key] = "student"
	// or
	// myMap[myStruct{Name: "Raj", Add: "vns"}] = "student"
	fmt.Println("after update values===", myMap)
	//check key exists or not
	if values, exists := myMap[key]; exists {
		fmt.Println("Found:", values)
	} else {
		fmt.Println("Not found")

	}

}

func MapWhereTheValueIsAStruct() {
	type myStruct struct {
		Name string
		Add  string
	}

	myMap := map[string]myStruct{
		"student":  {Name: "Ved", Add: "Varanbasi"},
		"Empolyee": {Name: "Ram", Add: "Ayodhya"},
	}

	// add new values
	myMap["company"] = myStruct{Name: "TCS", Add: "Gujarat"}

	fmt.Println("map where the value is a struct==", myMap)

	// modified the data
	myMap["student"] = myStruct{Name: "RAHUL VERMA", Add: "kardhana"}
	fmt.Println("updated map where the value is a struct==", myMap)
	// check key exists or not
	if values, exists := myMap["student"]; exists {
		fmt.Println("Found:", values)
	} else {
		fmt.Println("Not found")

	}
}

func ConvertAMapToAJSONstring() {
	myMap := map[string]string{
		"name": "Raghav",
		"add":  "lucknow",
	}
	jsonData, err := json.Marshal(myMap)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("convert a map to a JSON string", string(jsonData))
}

func ConvertAJSONstringToAMap() {
	jsonData := `{"Name":"ved","Add":"bodhapur","city":"varanasi","state":"up"}`

	var myMap = make(map[string]interface{})

	err := json.Unmarshal([]byte(jsonData), &myMap)
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	fmt.Println("convert a JSON string to a map", myMap)
}

func HandleConcurrentAccessToAMap() {

}

func FindTheKeyWithTheHighestValue() {
	myMap := map[string]int{
		"Alice":   50,
		"Bob":     85,
		"Charlie": 92,
		"David":   78,
	}

	var maxValue int
	var maxKey string

	for key, value := range myMap {
		if value > maxValue {
			maxValue = value
			maxKey = key
		}
	}
	fmt.Printf("Key with highest value: %s (Value: %d)\n", maxKey, maxValue)

}

func FindTheKeyWithTheLowestValue() {
	data := map[string]int{
		"Alice":   50,
		"Bob":     85,
		"Charlie": 92,
		"David":   28,
	}

	var minKey string
	minValue := data["Alice"]

	for key, values := range data {
		if values < minValue {
			minValue = values
			minKey = key
		}
	}
	fmt.Printf("key with the lowest values: %s (values %d)\n", minKey, minValue)
}

func RemoveMultipleKeysFromAMap() {
	data := map[string]int{
		"Alice":   50,
		"Bob":     85,
		"Charlie": 92,
		"David":   78,
		"Eve":     67,
	}

	mySlice := []string{"Alice", "Bob"}

	for _, key := range mySlice {
		delete(data, key)
	}
	fmt.Println("Updated map:", data)

}

func CountOccurrencesOfElementsUsingAMap() {
	words := []string{"apple", "banana", "apple", "orange", "banana", "apple"}

	myMap := map[string]int{}

	for _, word := range words {
		myMap[word]++
	}
	fmt.Println("Word counts:", myMap)
}

var cache = make(map[int]string)

func MapAsACache() {
	fmt.Println(getData(1)) // Fetches from DB
	fmt.Println(getData(2)) // Fetches from DB
	fmt.Println(getData(1)) // Retrieves from cache
}

func getData(id int) string {
	if value, exists := cache[id]; exists {
		fmt.Println("Cache hit!")
		return value
	}
	// Simulate fetching from DB (if not in cache)
	fmt.Println("Fetching from DB...")
	data := fmt.Sprintf("User:-%d", id)

	cache[id] = data
	return data
}

func MapWhereTheValueIsASlice() {
	// Initializing a map
	myMap := map[string][]string{
		"Student": {"ved", "rahul"},
	}
	fmt.Println("data of map", myMap)
	// Appending a new subject for Ved
	myMap["Student"] = append(myMap["Student"], "Ram", "Rohan")
	// Adding subjects for a student
	myMap["Teacher"] = []string{"Ram", "Rohan"}
	fmt.Println("after add more data:==", myMap)
}

func IteratingOverAMapOfSlices() {
	myMap := map[string][]int{
		"varanasi": {22221309, 229089, 980786, 909876},
		"Lucknow":  {698098, 359028, 984623},
	}

	for city, zipCode := range myMap {
		fmt.Println("city:", city, "zip Code:", zipCode)
	}
}
