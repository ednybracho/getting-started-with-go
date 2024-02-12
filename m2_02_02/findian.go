/*
Write a program which prompts the user to enter a string. The program searches through the entered string for the characters ‘i’, ‘a’, and ‘n’. The program should print “Found!” if the entered string starts with the character ‘i’, ends with the character ‘n’, and contains the character ‘a’. The program should print “Not Found!” otherwise. The program should not be case-sensitive, so it does not matter if the characters are upper-case or lower-case.

Examples: The program should print “Found!” for the following example entered strings, “ian”, “Ian”, “iuiygaygn”, “I d skd a efju N”. The program should print “Not Found!” for the following strings, “ihhhhhn”, “ina”, “xian”.

Submit your source code for the program,
“findian.go”.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var str string
	fmt.Print("Please introduce a text:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	str = scanner.Text()
	strLower := strings.ToLower(str)
	initial := strings.HasPrefix(strLower,"i")
	final := strings.HasSuffix(strLower,"n")
	inTheMiddle := strings.Index(strLower,"a")
	if (initial && final && (inTheMiddle>0 && inTheMiddle<len(strLower)-1)) {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}

	avg := 2 % (10 + 10)
	fmt.Println(avg)

	i, _ := strconv.Atoi("10")
  y := i * 2
  fmt.Println(y)

	s := strings.Replace("ianianian", "ni", "in", 2)
  fmt.Println(s)

	x:=7
  switch {
    case x>3:
      fmt.Printf("1")
    case x>5:
      fmt.Printf("2")
    case x==7:
      fmt.Printf("3")
    default: 
      fmt.Printf("4")
  }

	var xtemp int
  x1 := 0
  x2 := 1
  for x:=0; x<5; x++ {
    xtemp = x2     // 1 - 1 - 2 - 3 - 5
    x2 = x2 + x1   // 1 - 2 - 3 - 5 - 8
    x1 = xtemp     // 1 - 1 - 2 - 3 - 5
  }
  fmt.Println(x2)

	var xed int
  var yed *int
  zed := 3
  yed = &zed
  xed = &yed
}