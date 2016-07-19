# Go TDD Example - Factorials

An example of Test-Driven Development in Go against a function that generates factorials.

E.g. 4! = 4 * 3 * 2 * 1 = 24


## Running Locally

* `git clone https://github.com/jamesseanwright/golang-tdd-factorial.git` to `$GOPATH/src`
* `cd golang-tdd-factorial`
* `go test -v ./...`

## Commits

### `git checkout master~2`

Contains tests and an unimplemented `GetFactorial` function, that returns -1 when invoked.


### `git checkout master~1`

Implements the `GetFactorial` function with a `for` loop that satisfies the tests.


### `git checkout master`

Refactors the `GetFactorial` function to use recursion. The tests still pass.