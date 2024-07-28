from observer import Store, Customer


store = Store()

customer1 = Customer("Customer 1")
customer2 = Customer("Customer 2")

store.add_subscriber(customer1)
store.add_subscriber(customer2)

store.notify()
# Outputs:
# Customer 1 has been notified.
# Customer 2 has been notified.

store.remove_subscriber(customer1)

store.notify()
# Outputs:
# Customer 2 has been notified.
