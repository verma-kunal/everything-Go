# Structs, Methods, Interfaces

- [method](https://go.dev/ref/spec#Method_declarations) - "not the traditional function". it's a function with a "receiver"
    > Methods are very similar to functions but they are called by invoking them on an instance of a particular type.
- interface - In Go, interfaces are a way to define behavior. They specify a set of method signatures, but unlike other languages, they don’t provide the actual implementation. "Any type that implements all the methods of an interface is considered to implement that interface".
    > This allows Go to achieve polymorphism!
- [table driven tests](https://go.dev/wiki/TableDrivenTests) - useful when you want to build a list of test cases that can be tested in the same manner.


