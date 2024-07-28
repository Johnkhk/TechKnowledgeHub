package main

import "fmt"

// Subscriber interface that all subscribers must implement
type Subscriber interface {
	Update()
}

// Customer is a concrete subscriber that implements the Subscriber interface
type Customer struct {
	name string
}

func (c *Customer) Update() {
	fmt.Printf("%s has been notified.\n", c.name)
}

// Store is the publisher that maintains a list of subscribers and notifies them
type Store struct {
	subscribers []Subscriber
}

func (s *Store) addSubscriber(subscriber Subscriber) {
	s.subscribers = append(s.subscribers, subscriber)
}

func (s *Store) removeSubscriber(subscriber Subscriber) {
	for i, sub := range s.subscribers {
		if sub == subscriber {
			s.subscribers = append(s.subscribers[:i], s.subscribers[i+1:]...)
			break
		}
	}
}

func (s *Store) notify() {
	for _, subscriber := range s.subscribers {
		subscriber.Update()
	}
}
