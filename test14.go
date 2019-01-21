package main

import "fmt"

// map

func main() {
	map1 := make(map[string]string)
	map1["a"] = "hello "
	map1["b"] = "world"
	fmt.Println(map1)
	delete(map1, "a")
	fmt.Println(map1)

	var map2 map[string]string
	map2 = make(map[string]string)
	map2["a"] = "test"
	fmt.Println(map2)

	map3 := map[string]string{}
	map3["a"] = "test3"
	fmt.Println(map3)

	var map4 = map[string]string{}
	map4["a"] = "test4"
	fmt.Println(map4)
}