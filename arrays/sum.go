/*
https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/arrays-and-slices

What we learn:

- Arrays
- Slices
  -The various ways to make them
  -How they have a fixed capacity but you can create new slices from old ones using append
  -How to slice, slices!
- len to get the length of an array or slice
- Test coverage tool
- reflect.DeepEqual and why it's useful but can reduce the type-safety of your code


go test -cover to see the coverage of your tests
*/

package arrays

//Sum the numbers passed by parameters
func Sum(numbers []int) int {
	sum := 0

	for _, number := range numbers {
		sum += number
	}

	return sum
}

//SumAll slices passed by parameter
func SumAll(numbersToSum ...[]int) []int {
	var sums []int

	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}

//SumAllTails sum all numbers on the tail of the slice passed
func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int

	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:] //slice[low:high]
			sums = append(sums, Sum(tail))
		}

	}

	return sums
}
