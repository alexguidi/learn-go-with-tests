/*
https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/integers

What we learn:
- More practice of the TDD workflow
- Integers, addition
- Writing better documentation so users of our code can understand its usage quickly
- Examples of how to use our code, which are checked as part of our tests*/

package integers

//Add takes two integers and returns the sum of them.
func Add(x, y int) int {
	return x + y
}
