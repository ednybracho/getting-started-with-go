/*
Write a program which reads information from a file and represents it in a slice of structs. Assume that there is a text file which contains a series of names. Each line of the text file has a first name and a last name, in that order, separated by a single space on the line.

Your program will define a name struct which has two fields, fname for the first name, and lname for the last name. Each field will be a string of size 20 (characters).

* Your program should prompt the user for the name of the text file.
* Your program will successively read each line of the text file and create a struct which contains the first and last names found in the file.
* Each struct created will be added to a slice, and after all lines have been read from the file, your program will have a slice containing one struct for each line in the file.
* After reading all lines from the file, your program should iterate through your slice of structs and print the first and last names found in each struct.

! Submit your source code for the program, “read.go”.
*/
package main

import (
	"fmt"
	"os"
)

type Person struct {
	fname string
	lname string
}
func main () {
	fmt.Print("Please Could you enter name's file? ")
	var file string
	fmt.Scan(&file)
	var sliPerson = make([]Person,0,1)

	f, _ := os.Open(file)
	barr := make([]byte,1024)
	nb, _ := f.Read(barr)
	// fmt.Printf("Read %d bytes: %s\n", nb, barr[:nb])

	var fname string = ""
	var lname string = ""
	var isName bool = true
	for _, v := range barr[:nb] {
		if v == byte(32) {
			// fmt.Println("Space")
			isName = false
		} else if v == byte(10) {
			// fmt.Println("New Line")
			var newPerson = new(Person)
			newPerson.fname = fname
			newPerson.lname = lname
			sliPerson = append(sliPerson, *newPerson)
			fname = ""
			lname = ""
			isName = true
		} else {
			if isName {
				fname += string(v)
			} else {
				lname += string(v)
			}
		}
		// fmt.Printf("%d, %c",v,v)
	}
	if lname != "" {
		var newPerson = new(Person)
		newPerson.fname = fname
		newPerson.lname = lname
		sliPerson = append(sliPerson, *newPerson)
	}

	for _,v := range(sliPerson) {
		fmt.Printf("First name is %s, and Last name is %s \n", v.fname, v.lname)
	}

	f.Close()

}
