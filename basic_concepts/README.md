# Basics of Go

## Syntax 101

- project initialization: 
```go
go mod init example.com/intro
go mod tidy
```
- all the Go code belongs to a certain package. Everytime we create a `.go` file, we have to define the package being used at the top:
```go
package main
```
- the entrypoint of the main Go code i.e. from where the Go compiler will start the execution of the program, is a function. For example:
```go
func main(){}
```
- to import external to the file, we use the `import` keyword followed by the package name. And then, we can reference the package to use its functionalities (functions inside a package). For example:
```go
import fmt

// using the package
fmt.Println("hello")
```
- Run the Go program: 
```go
go run main.go
```

## Variables

- Go compiler throws an error if we define a variable but not use it - To enforce better code quality
- We can use `fmt.Printf` to format values in a print statement. For example:
```go
var confName = "GopherCon Community Edition 2024"
fmt.Printf("Welcome to %v!\n", confName)
```
- As a shortcut, we can use `:=` to declare and assign values to a variable. Called as "short variable declaration".
  - Doesn't work with `const` and if we want to explicitly define a variable type

## Data Types

- Go is statically typed language:
  - We need to tell the data type when declaring a new variable, or
  - Type Inference: Go can infer the type automatically - when a value is assigned (according to that value)
- Explicitly defining the variable type:
```go
var userName string
var userTickets int

userName = "Toby"
userTickets = 3
```

## Pointers

- A special variable that references another variable's value's address in memory.
- Explanation in Hindi:
```text
Jab tum ek variable declare karte ho, toh us variable ka ek memory address hota hai. 
Pointer ek aisa variable hota hai jo dusre variable ka wohi address store karta hai.
```
- For example:
```go
 // Declare an integer variable
num := 10

// Declare a pointer variable
var pointer *int

// Assign the address of the variable to the pointer
pointer = &num

// Print the value of the variable
fmt.Println("Value of num:", num)

// Print the address of the variable
fmt.Println("Address of num:", &num)

// Print the value stored in the pointer
fmt.Println("Value stored in pointer:", *pointer)
```
## User Input

```go
var firstName string
fmt.Print("Enter your first name:")
fmt.Scan(&firstName) // using a pointer here
```

## Arrays & Slices

### Arrays

- arrays in Go have fixed size i.e. how many elements it can hold
- declaring an empty array:
```go
// 1
var bookings = [50]string{}

// 2 - declaring an array type variable:
var bookings [50]string
```

### Slice

- abstraction of an array i.e. works as arrays under-the-hood but has a variable length
- index-based and have a size (can be re-sized)
- basic syntax of an empty slice:
```go
var bookings []string

// appending values in the slice:
bookings = append(bookings, firstName+" "+lastName)
```

### Loops
- Loops are simplified in Go - we only have a 'for loop', but it has different types as well:
  - infinite loop
    ```go
    for {}
    ```
  - for-each loop
    - here, a 'range' keyword helps us to iterate over an entire data type (not just array or slice) and returns the index & the element itself
    - `_` is called a blank identifier that helps to ignore the vars that we don't use. For example:
      ```go
      for _, booking := range bookings {
			names := strings.Fields(booking) // converts the string name into a slice (elements separated by ',')

			// adding the firstname to our new slice
			firstNames = append(firstNames, names[0])
		}
      ```
    
### Conditionals

- Example (if):
```go
if remainTickets == 0 {
  // end program
  fmt.Println("Tickets are sold out! See you next year.")
  break
}
```
- `continue` - skips the rest of the body in the loop & directly moves to the next iteration.
- Different types:
  - `if`
  - `if...else`
  - `else...if` - can have as many b/w an `if...else` block

### Functions

- aim: making the code cleaner & descriptive
- normally, we cannot return multiple values in a function, but in Go we can:
```go
func validateUserInput(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {

	// input validation
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@") // gives back boolean result
	isValidTicketNum := userTickets > 0 && userTickets < remainTickets

	return isValidName, isValidEmail, isValidTicketNum
}
```

### Package Level Variables

- defined outside all the functions & are accessible by all of them

### Go Packages

- package: a collection of Go files 
- variables & functions defined outside a function, can be accessed in all other files 'within the same package'
- multiple packages. A few pointers while using multiple packages:
  - the particular needs to be explicitly imported to files outside the package - for functions to be used
    ```go
    import (
	helper "Go-Learn/basics/helper" // "module_name/package"
	"fmt"
	"strings"
	)
    
    // referencing the function from a file in helper package
    helper.ValidateUserInput(firstName, lastName, email, userTickets)
    ```
  - exporting a variable or function - start the name with a CAPITAL letter:
    ```go
    func ValidateUserInput(firstName string, .....
    ```
- Variables: 3 Levels of Scope
  1. Local - declared within functions or blocks (like if..else). can be used only within functions or those blocks
  2. Package - direct access to the entire package
  3. Global

### Maps

- a data type that allows to store data in unique key-value pairs
- empty map syntax:
  ```go
  userData := make(map[string]string) // "map[key_data_type]value_data_type"
  
  // adding values:
  userData["firstName"] = firstName
  userData["lastName"] = lastName
  userData["email"] = email
  userData["numOfTickets"] = strconv.FormatUint(uint64(userTickets), 10) // converting 'uint' to 'string'
  ```
  
### Struct

- key-value data structure that can store mix data types (unlike a map)
- basic syntax:
  ```go
  type UserData struct {
	    firstName    string
	    lastName     string
	    email        string
	    numOfTickets uint
  }
  ```
- basically, we can create a custom type and give it a new "value-type" it can take i.e. properties.

### Goroutines - Concurrency Concept

- goroutine - a lightweight thread managed by the Go runtime
- using a `go` keyword: starts the execution of a function call as an independent concurrent thread of control, or goroutine, within the same address space.
  ```go
  go sendTicket(userTickets, firstName, lastName, email)
  ```
- by default, the main goroutine (or thread) does not wait for other goroutines. Therefore, if a function in the main thread exits i.e. the entire program exits in the main thread, rest of the threads will not even execute!
  - Solution: Make the main goroutine wait, till other threads get executed (the ones getting created separately)
    - using a `WaitGroup` - waits for the launched (main) goroutine to finish
      - basic syntax:
        ```go
          var wg = sync.WaitGroup{} // global declaration
        ```
      - different functions we use: `Add()`, `Wait()`, `Done()`
- Communication between different gorountines can be done using [channels](https://www.educative.io/answers/what-are-channels-in-golang) in Go

## Interesting Bits

- Difference b/w `Printf` and `Sprintf`:
The `fmt.Sprintf` function is used to format strings and return the resulting string without printing it to the standard output.
It is similar to `fmt.Printf`, but instead of printing the formatted string, it returns the string as a value.

# References/Resources

- [Beginner's Tutorial by TechWorld with Nana](https://youtu.be/yyUHQIec83I?feature=shared)
- Refer the [standard library](https://pkg.go.dev/std) for more Go packages
- [strconv package](https://pkg.go.dev/strconv): conversions to and from string representations of basic data types