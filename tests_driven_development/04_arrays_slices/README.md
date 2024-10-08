# Arrays & Slices

- The size of array is encoded in its type. If you try to pass an `[4]int` into a function that expects `[5]int`, it won't compile - they are of different types!
- slices: have "variable" length
- [Go's in-built coverage tool](https://go.dev/blog/cover): the coverage tool can help identify areas of your code not covered by tests.

  ```bash
  $ go test -cover

  PASS
  coverage: 100.0% of statements
  ok      github.com/verma-kunal/everything-Go/tests_driven_development/04_arrays_slices  0.395s
  ```

- [variadic functions](https://gobyexample.com/variadic-functions): can be called with any number of trailing arguments.
  ```go
  func sum(nums ...int){}
  ```
- `reflect.DeepEqual()`: can be used to compare to variables (particularly slices).
  - (important) - is not "type safe". therefore, be careful using while using it.
    > From Go `1.21`, slices standard package is available, which has [`slices.Equal`](https://pkg.go.dev/slices#Equal) function to do a simple shallow compare on slices.
- 
