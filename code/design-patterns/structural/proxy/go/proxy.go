package main

import (
	"fmt"
)

// ServiceInterface represents the interface for the service
type ServiceInterface interface {
	handleRequest(url string) string
}

// WebServer represents the actual service
type WebServer struct{}

// handleRequest handles the request directly in the WebServer
func (ws *WebServer) handleRequest(url string) string {
	return fmt.Sprintf("WebServer handling request for: %s", url)
}

// Nginx represents the proxy for the WebServer
type Nginx struct {
	webServer   *WebServer
	rateLimiter map[string]int
	cache       map[string]string
	maxRequests int
}

// NewNginx creates a new Nginx proxy with a given maxRequests limit
func NewNginx(maxRequests int) *Nginx {
	return &Nginx{
		webServer:   &WebServer{},
		rateLimiter: make(map[string]int),
		cache:       make(map[string]string),
		maxRequests: maxRequests,
	}
}

// handleRequest controls access, rate limits, and caches requests
func (nginx *Nginx) handleRequest(url string) string {
	if !nginx.controlAccess(url) {
		return "Access denied"
	}

	if response, found := nginx.cache[url]; found {
		fmt.Println(response)
		return fmt.Sprintf("Returning cached response for: %s", url)
	}

	response := nginx.webServer.handleRequest(url)
	nginx.cacheRequest(url, response)
	return response
}

// controlAccess checks the rate limit for the given url
func (nginx *Nginx) controlAccess(url string) bool {
	if count, exists := nginx.rateLimiter[url]; exists {
		if count >= nginx.maxRequests {
			return false
		}
		nginx.rateLimiter[url]++
	} else {
		nginx.rateLimiter[url] = 1
	}
	return true
}

// cacheRequest caches the response for the given url
func (nginx *Nginx) cacheRequest(url, response string) {
	nginx.cache[url] = response
}
