package main

import (
	"fmt"
	"time"
)

// simulate requests
func simulateRequests(nginx *Nginx, urls []string) {
	for _, url := range urls {
		fmt.Println(nginx.handleRequest(url))
		time.Sleep(100 * time.Millisecond) // simulate delay between requests
	}
}

func main() {
	nginx := NewNginx(2)

	urls := []string{
		"/home",
		"/home",
		"/home", // should be denied
		"/about",
		"/home", // should be denied
		"/home", // should be denied
		"/contact",
	}

	simulateRequests(nginx, urls)
}
