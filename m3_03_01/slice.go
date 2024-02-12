/*
? Write a program which prompts the user to enter integers and stores the integers in a sorted slice.
  - The program should be written as a loop.
  - Before entering the loop, the program should create an empty integer slice of size (length) 3.
  - During each pass through the loop, the program prompts the user to enter an integer to be added to the slice.
  - The program adds the integer to the slice, sorts the slice, and prints the contents of the slice in sorted order.
  - The slice must grow in size to accommodate any number of integers which the user decides to enter.
  - The program should only quit (exiting the loop) when the user enters the character ‘X’ instead of an integer.

! Submit your source code for the program, “slice.go”.
*/
package main

import (
	"fmt"
)

type P struct {
	x string
y int
}

func main() {
	// fmt.Println("Start the program!")
	// var sliIntegers = make([]int,3)
	// var pos int = 0
	// for true {
	// 	fmt.Print("Please introduce an Integer value or 'X' to quit: ")
	// 	var num int
	// 	_, err:= fmt.Scan(&num)
	// 	if err != nil {
	// 		break
	// 	}
	// 	if pos < 3 {
	// 		sliIntegers[pos] = num
	// 		pos++
	// 	} else {
	// 		sliIntegers = append(sliIntegers, num)
	// 	}
	// }
	// var order bool = true
	// var iteration int  = 0
	// for order  {
	// 	order = false
	// 	iteration++
	// 	var long = len(sliIntegers) - iteration
	// 	for i:=0; i < long; i++{
	// 		if sliIntegers[i] > sliIntegers[i+1] {
	// 			order = true
	// 			var aux = sliIntegers[i]
	// 			sliIntegers[i] = sliIntegers[i+1]
	// 			sliIntegers[i+1] = aux
	// 		}
	// 	} 
	// }
	// fmt.Println(sliIntegers)

	// fmt.Println("Finished the program!")


	// x := []int {4, 8, 5}
  // y := -1
  // for _, elt := range x {
  //   if elt > y {
  //     y = elt
  //   }
  // }
  // fmt.Print(y)


	// x := [...]int {4, 8, 5}
  // y := x[0:2]
  // z := x[1:3]
  // y[0] = 1
  // z[1] = 3
  // fmt.Print(x)

	// x := [...]int {1, 2, 3, 4, 5}
  // y := x[0:2]
  // z := x[1:4]
  // fmt.Print(len(y), cap(y), len(z), cap(z))

	// x := map[string]int {
  //   "ian": 1, "harris": 2}
  // for i, j := range x {
  //   if i == "harris" {
  //     fmt.Print(i, j)
  //   }
 	// }

	//  b := P{"x", -1}
	//  a := [...]P{P{"a", 10}, 
	// 			 P{"b", 2},
	// 			 P{"c", 3}}
	//  for _, z := range a {
	// 	 if z.y > b.y {
	// 		 b = z
	// 	 }
	//  }
	//  fmt.Println(b.x)

	s := make([]int, 0, 3)
  s = append(s, 100)
  fmt.Println(len(s), cap(s))
}