package main

func main() {
	store := &Store{}

	customer1 := &Customer{name: "Customer 1"}
	customer2 := &Customer{name: "Customer 2"}

	store.addSubscriber(customer1)
	store.addSubscriber(customer2)

	store.notify()
	// Outputs:
	// Customer 1 has been notified.
	// Customer 2 has been notified.

	store.removeSubscriber(customer1)

	store.notify()
	// Outputs:
	// Customer 2 has been notified.
}
