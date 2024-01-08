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

# References/Resources

- Refer the [standard library](https://pkg.go.dev/std) for more Go packages