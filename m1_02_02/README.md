# Workspaces

In this exercise we will talking about the structure of project in GO.

## Main Folders

- There are three subdirectories
  - src - Contains source code files
  - pkg - Contains packages (libraries)
  - bin - Contains executables

## Notes

- A lot of programmers have in the same workspaces differents projects, it's common to see that kind of situations.
- Directory hierarchy is recommended, not enforced.
  - i.e. you can have an executable in src
- Workspace directory defined by **GOPATH** enviroment variable
- GOPATH id defined during installation
  - C:\Users\yourname\go
- Go tools assume that code is in GOPATH

## Packages

- Group of related source files
- Each package can be imported by other packages
- Enables software reuses
- First line of file names the package
  - `package annpkg` - one file
  - `package billpkg` - another file
  - `import ( "annpkg" "billpkg" )` - another piece of code
- There must be one package called **main**
- Building the main package generated an executable program
- Main package needs a main() function
- main() is where code execution starts
- Example code:

```
package main
import "fmt"
func main() {
  ftm.println("hello, world\n")
}
```

## Import

- import keyword is used to access other packages
- Go standard library includes many packages
  - i.e "fmt"
- Searches directories specified by GOROOT and GOPATH

## The GO Tool

- a tool to manage Go source code
- Several commands:
  - _go build_ - compiles the program
    - Arguments can be a list of packages or a list of .go files
    - Creates and executable for the main package, same name as the first .go file
    - .exe suffix for executable in Windows
  - _go doc_
  - _go fmt_
  - _got get_ - downloads packages and installs them
  - _go list_
  - _go run_ - compiles .go files and runs the executable
  - _go test_ - runs tests using files ending in "_test.go"

## Variables
- Names are needed to refer to programming constructs
- Names must start with a letter
- Any number of letters, digits, underscores
- Case sensite
- Don't use keywords
- Must have a name and a type
  - ``` var x int```
  - ``` var x, y int```
- Types of variables
  - Integer
  - Floating point
  - Strings
    - Unicode
  

