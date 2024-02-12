/*
? Write a program which prompts the user to first enter a name, and then enter an address.
* Your program should create a map and add the name and address to the map using the keys “name” and “address”, respectively.
* Your program should use Marshal() to create a JSON object from the map, and then your program should print the JSON object.

! Submit your source code for the program, “makejson.go”.
*/

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	personMap := make(map[string]string)
	fmt.Print("Please Could you give your name? ")
	var name string
	fmt.Scan(&name)
	var address string
	fmt.Print("Please Could you give your address? ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	address = scanner.Text()
	personMap["name"] = name
	personMap["address"] = address
	barr, err := json.Marshal(personMap)
	if err != nil {
		fmt.Print("Error has happened!")
		}
	fmt.Println(string(barr))
}