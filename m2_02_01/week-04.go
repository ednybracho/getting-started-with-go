package main

/*

	! Protocols and Formats
	? Request for Comments (RFC)
		* Definition of Internet protocols and formats
		*  Examples
			* HTML -> Hypertext Markup Language, 1866
			* URI, Uniform resource identifier, 3986
			* HTTP -> Hypertext tranfer protocol, 2616
	? Protocol Packages
	* Golang provides packages for important RFCs
	* Functions which encode and decode protocol format
	* Example
		* "net/http"
			* Web comunication protocol
			* 	http.Get(www.uci.edu)
		* "net"
			* TCP/IP and socket programming
				* net.Dial("tcp","uci.edu:80")

	!JSON
	? JavaScript Object Notation
		* RFC 7159
		* Format to represent structured information
		* Atribute-value pairs
			* Struct or map
		* Basic value types
			* Bool, number, string, array, "object"
	? JSON Example
	* Go Struct
	p1 := Person(name: "joe", addr: "a st.", phone: "123")
	* Equivalent JSON Object
	{"name":"joe", "addr":"a st.", "phone":"123"}

	? JSON Properties
	* All Unicode
	* Human-readable
	* Fairly compact representation
	* Types can be combined recursively
		* Array of Structs, struct in struct

	? JSON Marshalling
	* Generating JSON representation from a object
	type struct Person {
		name string
		addr string
		phone string
	}

	p1 := Person(name:"joe",addr:"a st.", phone: "123")
	barr, err := json.Marshal(p1)
	* Marshal() return JSON representation as []byte

	? JSON Unmarshalling
	var p2 Person
	err := json.Unmarshal(barr,&p2)
	* Unmarshal() converts a JSON []byte into a Go Object
	* Pointer to Go Objects is passed to Unmarshal()
	* Object must "fit" JSON []byte

	! Files
	? Linear access, not random access
	* Mechanical delay
	? Basic operations
	* Open - get handle for access
	* Read - read bytes into []byte
	* Write - write []byte into file
	* Close - release handle
	* Seek - move read/write head

	? ioutil File Read
	* "io/ioutil" package has basic functions
		dat, e := ioutil.ReadFile("test.txt")
		* dat is []byte filled with contents of entire file
		* Explicit open/close at not needed
		* Large files cause a problem

		? ioutil File Write
			dat = "Hello, World!"
			err := ioutil.WriteFile("outfile.txt",dat,0777)
			* Writes []byte to file
			* Creates a file
			* Unix-style permission bytes

		! os Package File Access
		* os.Open() opens a file
			* Return a file descriptor (File)
		* os.Close() closes a file
		* os.Read() reads from a file into a []byte
			* Fills the []byte
			* Control the amount read
		* os.Write() writes a []byte into a file

		? os File Reading
		* Opening and reading
		f, err := os.Open("dt.txt")
		barr := make([]byte,10)
		nb, err := f.Read(barr)
		f.Close()
		* Reads and fills barr
		* Read returns # of bytes read
		* May be less than []bytes length

		? os File Create/Write
		f, err := is.Create("ooutfile.txt")
		barr := []byte{1,2,3}
		nb, err := f.Write(barr)
		nb, err := f.WriteString("Hi")
		* WriteString() writes a string
		* Write() writes a []byte
			* Any Unicode sequence



*/
func ProtocolsFormats() {



}