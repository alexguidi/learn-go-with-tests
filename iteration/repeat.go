/*
https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/iteration

What we learn:
- More TDD practice
- Learned for
- Learned how to write benchmarks
*/

package iteration

//Repeat a character passed
func Repeat(character string, times int) string {
	var repeated string

	for i := 0; i < times; i++ {
		repeated += character
	}

	return repeated
}
