/*
Write a program which prompts the user to enter a floating point number and prints the integer which is a truncated version of the floating point number that was entered. Truncation is the process of removing the digits to the right of the decimal place.

Submit your source code for the program, “trunc.go”.
*/

package main

import (
	"fmt"
)

func main() {
	var fNumber float32
	fmt.Printf("Please enter a floating Number:")
	num, err := fmt.Scan(&fNumber)
	if num == 1 {
		iNumber := int(fNumber)
		fmt.Printf("%d\n",iNumber)
	} else {
		if err != nil {  fmt.Printf("You don't enter a Number")}
	}
	
	

}