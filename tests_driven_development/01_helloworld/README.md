# Basics - Hello World with Tests

- good practice to separate the "domain" (core logic) & the "side-effects" (external systems or environments).
  - why?
    - pure domain logic (without side effects) is easier to test because it doesn't rely on external systems.
    - more reusability - the "logic" can be reused for different purpose.
- you can define "subsets" using `t.Run()` to test different scenarios for a particular piece of code.
- a typical cycle to follow:
  - Write a test
  - Make the compiler pass
  - Run the test, see that it fails and check the error message is meaningful
  - Write enough code to make the test pass
  - Refactor
- Basic TDD process:
  - Write a failing test and see it fail so we know we have written a relevant test for our requirements and seen that it produces an easy to understand description of the failure
  - Writing the smallest amount of code to make it pass so we know we have working software
  - Then refactor, backed with the safety of our tests to ensure we have well-crafted code that is easy to work with
