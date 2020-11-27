/*
https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/select

What we learned:

select
Helps you wait on multiple channels.
Sometimes you'll want to include time.After in one of your cases to prevent your system blocking forever.


httptest
A convenient way of creating test servers so you can have reliable and controllable tests.
Using the same interfaces as the "real" net/http servers which is consistent and less for you to learn.
*/

package main

import (
	"fmt"
	"net/http"
	"time"
)

var tenSecondTimeout = 10 * time.Second

//Racer makes a racer between two URLs to see which one is faster
func Racer(a, b string) (winner string, error error) {
	return ConfigurableRacer(a, b, tenSecondTimeout)
}

//ConfigurableRacer sets a timeout value
func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, error error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}
