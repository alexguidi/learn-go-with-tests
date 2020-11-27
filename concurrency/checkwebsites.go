/*
https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/concurrency

Wrapping up
This exercise has been a little lighter on the TDD than usual. In a way we've been taking part in one long refactoring of the CheckWebsites function; the inputs and outputs never changed, it just got faster. But the tests we had in place, as well as the benchmark we wrote, allowed us to refactor CheckWebsites in a way that maintained confidence that the software was still working, while demonstrating that it had actually become faster.

In making it faster we learned about

- goroutines, the basic unit of concurrency in Go, which let us check more
than one website at the same time.
- anonymous functions, which we used to start each of the concurrent processes
that check websites.
- channels, to help organize and control the communication between the
different processes, allowing us to avoid a race condition bug.
- the race detector which helped us debug problems with concurrent code

*/

package concurrency

//WebsiteChecker is a type
type WebsiteChecker func(string) bool
type result struct {
	string
	bool
}

//CheckWebsites check if the sites are up
func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		go func(u string) {
			resultChannel <- result{u, wc(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		result := <-resultChannel
		results[result.string] = result.bool
	}

	return results
}
