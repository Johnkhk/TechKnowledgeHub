class Subscriber:
    def update(self):
        raise NotImplementedError("Subclass must implement abstract method")


class Customer(Subscriber):
    def __init__(self, name):
        self.name = name

    def update(self):
        print(f"{self.name} has been notified.")


class Store:
    def __init__(self):
        self.subscribers = []

    def add_subscriber(self, subscriber):
        self.subscribers.append(subscriber)

    def remove_subscriber(self, subscriber):
        self.subscribers.remove(subscriber)

    def notify(self):
        for subscriber in self.subscribers:
            subscriber.update()
