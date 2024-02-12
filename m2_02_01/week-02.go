package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)
func practice() {

	const NUMBER = 3.14
	const (
		x = 1
		name = "Edny"
	)

// iota
	type Grades int
	const (
		A Grades = iota
		B 
		C 
		D 
		E 
	)
	fmt.Println("Hello, World!")
	// Pointers
	var i int = 1
	var j int
	var ip *int

	ip = &i
	j = *ip

	/*
	* Here we use a PrintF method of fmt librari
	*/
	fmt.Printf("Value of i:%d, store in this address: %p, and J has the same value of i: %d, store in this address %p\n",i,ip,j,&j)

	var fNumber float32 = 3.14
	var iNumber int = int(fNumber)
	fmt.Printf("Value of Integer after make a conversion: %d\n", iNumber)
	var r rune;
	r = 32
	var a bool 
	a = unicode.IsDigit(r)
	a = unicode.IsSpace(r)
	a = unicode.IsLetter(r)
	a = unicode.IsLower(r)
	a = unicode.IsPunct(r)
	r = unicode.ToUpper(r)
	r = unicode.ToLower(r)

	strings.Index("edny","e")
	strings.Compare("Edny", "edny")
	strings.HasPrefix("edny", "ed")
	strings.Replace("edny","ed","ma",0)
	strings.ToLower("EDNY")
	strings.ToUpper("edny")
	strings.Trim("   edny    ","")
	// strings.Index("a")
	var nInt int
	var err error 
	nInt, err = strconv.Atoi("12")
	var str string = strconv.Itoa(10)
	
	var strFloat string = strconv.FormatFloat(3.14,0,2,4)
	strconv.ParseFloat("3.14",8)
	
	
	fmt.Printf("%b", a)

	fmt.Println(nInt, str, strFloat, err)

	/*
	** MyMethod
	? Control Flows
	* If Statement
	if <condition> {}
	if <condition> {} else {}
	if <condition> {} else if {} else {}
	
	@param loop
	* Traditional loop
	for <init>; <cond>; <update> {<stmts>}
	* Like a While Loop
	<init>
	for <cond> {
		<stmts>
		<update>
	}
	* Infinitive Loop
	for {} -> This is a infinite loop

	? Switch 
	* with tag
	switch <tag> {
		case 1: <stmt>
		case 2: <stmt>
		default: <stmt> 
	}

	* without tag(taglees)
	switch {
		case <cond> : <stmt>
		case <cond> : <stmt>
		default: <stmt>
	}

	? Break and Continue Statement
	* Break -> break the execution of the loop
	for i := 0; i< 10; i++ {
		if i == 5 { break }
		fmt.Println("Value")
	}

	* Continue -> Skip the next statements inside of the loop when the condition is met 
	for i := 0; i< 10; i++ {
		if i == 5 { continue }
		fmt.Println("Value")
	}

	? Scan
	* Scan is used to read user input
	* - Takes a pointer as an argument
	* - Typed data is written to pointer
	* - Returns number of scanned items
	* Example
	var appleNum int
	fmt.Printf("Number of apples?")
	num, err := fmt.Scan(&appleNum)
	fmt.Printf(appleNum)



	*/


}