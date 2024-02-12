/*

? Data Types

*/

package main

import "fmt"

func dataTypes () {
	fmt.Println("Starts with Data Types")

	/*
	! Arrays
	* Fixed-lenght series of elements of a chosen type
	* Elements accessed using subscript notatio, [ ]
	* Indices start at 0
	* Elements initialized to zero value
	var x [5]int
	x [0] = 2
	fmt.Printf(x[1])
	
	? Array Literal
	* An array pre-defined with values
	var x [5]int = [5]{1,2,3,4,5}
	* Length of literal must be length of array
	* ... for size in array infers size from numbers of initializers
	x:= [...]int{1,2,3,4}
	
	? Iterating through Arrays
	* User a for loop with the range keyword
	x:= [3]int {1,2,3}

	for i, v := range x {
		fmt.Printf(ind %d), val %d", i, v)
	}
	* Range returns two values
	* Index and element at index
	

	! Slices -> New on Go

	* A @window on an inderlying array
	* Variable size, up to the whole array
	* Pointer indicates the start of the slice
	* Length is the number of elements in the slice -> len()
	* Capacity is maximum number of elements -> cap()
		* from start of slice to end of array
	arr := [...]string{"a,"b","c","d","e","f","g"}
	s1 := arr[1:3]
	s2 := arr[2:5]
	
	? Accessing Slices
	? Slice Literals
	* Can be used to initialized a slice
	* Creates the underlying array and creates a slice to reference it
	* Slice points to the start of the array, length is capacity
	sli := []int{1,2,3}

	? Make
	* Create a slice (and array) using make()
	* 2-argument version: specifu type and length/capacity
	* Init. to zero, length = capacity
	sli = make([]int,10)
	* 3-argument version: specify length and capacity separately
	sli = make([]int, 10, 15)

	? Append
	* Size of a slice can be increased by append()
	* Adds elements to the end of a slice
	* Inserts into underlying array
	* Increases size of array if necessary
	sli = make([]int, 0, 3)
	* Length of sli is 0
	sli = append(sli,100)

	! Hash Tables
	* Contains key/value pairs
		* SSN/email
		* GPS coors/address
	* Each value is associated with a unique key
	* Hash function is used to compute the slot for a key

	? Tradeoffs of Hash Tables
	* Advantages
		* Faster lookup than lists
			* Constant-time, vs linear-time
		* Arbitrary keys
			* Not ints, like slices or arrays
	* Disadvantages
		* May have collisions
			* Two keys hash to same slot
		
	! Maps
	* Implementation of a Hash Table
	* Use make() tp create a map
	var idMap map[string]int
	idMap = make(map[string]int)
	* May define a map literal
	idMap := Map[string]int {"joe":123}
	
	? Accessing Maps 
	* Referencing a value with [key]
	* Returns zero if key is not present
	fmt.Println(idMap["joe"])  
	* Adding a key/value pair
	idMap["jane"] = 456
	* Deleting a key/value pair
	delete(idMap, "joe")

	? Map Functions
	* Two-value assigments tests for existence of the key
	id,p := idMap["joe"]
	* id is value, p is presence of key
	* len() returns number values
	fmt.Println(len(idMap))

	? Itrating through a Map
	* Use a for loop with the range keyword
	* Two-value assignment with range
	for key, val := range idMap {
		fmt.Println(key,val)
	}

	? Structs
	* Aggregate data type
	* Groups together other objects of arbitrary type

	* Example: Person struct
		* name, address, phone
		* option 1: have 3 separate variables, programmer remembers that they are related
		* option 2: Make a single struct which contains all 3 variables (vars)

	type struct Person {
		name string
		addr string
		phone string
	}

	var p1 person
	* each propertyis a field
	* p1 contains values for all fields

	? Accessing Struct fields
	* Use dot notation
	p1.name = "joe"
	x = p1.addr

	? Initializing Structs
	* Can use new()
	* Initializes field to zero
	p1 ;= new(Person)
	* Can initialize using a struct literal
	p1 := Person(name: "joe", addr:"a.st.", phone:"123")


	*/
}