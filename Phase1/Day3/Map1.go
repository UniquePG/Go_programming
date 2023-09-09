package main

import "fmt"

func main() {

	// map -> map is a datatype that stores in elements in the key value pairs
	//* syntax ->   var demo = map[keyType]valueType{ key : value, key : value, ....}
	//* 				demo := map[keyType]valueType{}
	// now add values in map ->
	// 							demo[key] = value

	//! 1st method to create a map
	var map1 = map[string]int{}			// key-> string, value-> int

	map1["aaaa"] = 22
	map1["bbbb"] = 55

	fmt.Println("access value of map1", map1["aaaa"])
	fmt.Println("access value of map1", map1["bbbb"])

	//! 2st method to create a map
	map2 := map[int]string{
		1: "my",
		2: "name",
		3: "is",
		4: "pavan", 	// , is neccessory at the end of the last element

	}



	fmt.Println("access value of map2", map2[1])
	fmt.Println("access value of map2", map2[2])
	fmt.Println("access value of map2", map2[3])
	fmt.Println("access value of map2", map2[4])

	// access map values using for loop
	for i := 0; i <= len(map2); i++ {				// len() func. is used to find the length
		fmt.Println("using for loop: ", map2[i])
	}

	// insert and update an element in the map
	map2[5] = "gupta"
	fmt.Println("insert a value in the map: ", map2[5])
	

	//* to delete any element in the map
	delete(map2, 5)				 //* delete(mapName, keyName)

	fmt.Println("after deleting an element at index 5: ", map2[5]);


	// if want to check value exists in the map or not

	valueIfExists, isAvailable := map2[4];		// value at index 4 -> pavan, true  
	valueIfExists2, isAvailable2 := map2[6];		// value at index 6 -> "", false
	

	fmt.Println("if value(key) exists: ", valueIfExists, isAvailable)
	fmt.Println("if value(key) not exists: ", valueIfExists2, isAvailable2)

		// we can omit if dont want
		_, available := map2[2];			// key 2 exist so it will return true

		fmt.Println("omit value in shorhand: ", available)

}
