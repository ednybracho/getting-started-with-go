package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Person struct {
	fname string
	lname string
}

func main() {
	var file_name string
	fmt.Scan(&file_name)
	var people []Person
	file, err := os.Open(file_name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := strings.Fields(scanner.Text())
		people = append(people, Person{fname: s[0], lname: s[1]})
	}

	for _, data := range people {
		fmt.Println(data.fname, data.lname)
	}
}
