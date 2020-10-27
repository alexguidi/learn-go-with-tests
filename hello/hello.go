/*
https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/hello-world

What we learn:
- Writing tests
- Declaring functions, with arguments and return types
- if, const and switch
- Declaring variables and constants


The TDD process and why the steps are important
Write a failing test and see it fail so we know we have written a relevant test for our requirements and seen that it produces an easy to understand description of the failure
Writing the smallest amount of code to make it pass so we know we have working software
Then refactor, backed with the safety of our tests to ensure we have well-crafted code that is easy to work with
*/

package main

import (
	"fmt"
)

const spanish = "Spanish"
const french = "French"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

//Hello func
func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("world", ""))
}
