# Notes

- use [benchmarks](https://pkg.go.dev/testing#hdr-Benchmarks) to test the performace of code.

  ```bash
  $ go test -bench=.

  goos: darwin
  goarch: arm64
  pkg: github.com/verma-kunal/everything-Go/tests_driven_development/03_iteration
  BenchmarkRepeat-11      16914445                72.09 ns/op
  PASS
  ok      github.com/verma-kunal/everything-Go/tests_driven_development/03_iteration      2.531s
  ```

  here:

  - `BenchmarkAdd-11`: This shows the name of the benchmark and the number of CPU cores used.
  - `16914445`: The number of iterations (how many times the code was run).
  - `72.09 ns/op`: The time taken per operation, in nanoseconds.
