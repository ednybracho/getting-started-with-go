package main

import ("fmt"
	"os"
	"bufio"
	"strings")



type Person struct{
	fname  string
	lname string
}

func main(){

	var filename string
	var slicePersons = make([]Person, 0, 1)

	fmt.Printf("Enter filename:")
	fmt.Scanf("%s", &filename)

	file, err := os.Open(filename)	
	if err != nil{
		fmt.Printf("Error: File %s wasn't found", filename)
		file.Close()
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	var line string
	for{
		line, err = reader.ReadString('\n')
		if err != nil{
			break
		}
		line = strings.TrimSuffix(line, "\n")
		personData:= strings.Split(line, " ")
		if len(personData) < 2 {
			continue
		}
		person := Person{fname: personData[0], lname: personData[1]}
		slicePersons = append(slicePersons, person)
	}
	file.Close()
	for _, person := range slicePersons{
		fmt.Println(person.fname, person.lname)
	}
}
