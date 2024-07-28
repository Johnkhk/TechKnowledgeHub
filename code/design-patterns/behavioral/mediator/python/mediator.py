from abc import ABC, abstractmethod
from typing import List


class Mediator(ABC):
    @abstractmethod
    def can_arrive(self, train: "Train") -> bool:
        pass

    @abstractmethod
    def notify_about_departure(self):
        pass


class StationManager(Mediator):
    def __init__(self):
        self.train_queue: List[Train] = []

    def can_arrive(self, train: "Train") -> bool:
        if not self.train_queue:
            self.train_queue.append(train)
            return True
        else:
            print(f"{train.__class__.__name__} is waiting for the platform.")
            return False

    def notify_about_departure(self):
        if self.train_queue:
            departed_train = self.train_queue.pop(0)
            print(f"{departed_train.__class__.__name__} has departed.")
        if self.train_queue:
            next_train = self.train_queue[0]
            print(f"{next_train.__class__.__name__} can now arrive.")
            next_train.arrive()


class Train(ABC):
    def __init__(self, mediator: Mediator):
        self.mediator = mediator

    @abstractmethod
    def arrive(self):
        pass

    @abstractmethod
    def depart(self):
        pass


# Client code is simple, just ask the mediator
class PassengerTrain(Train):
    def arrive(self):
        if self.mediator.can_arrive(self):
            print("Passenger train arrived.")
        else:
            print("Passenger train waiting for arrival.")

    def depart(self):
        print("Passenger train leaving...")
        self.mediator.notify_about_departure()


class FreightTrain(Train):
    def arrive(self):
        if self.mediator.can_arrive(self):
            print("Freight train arrived.")
        else:
            print("Freight train waiting for arrival.")

    def depart(self):
        print("Freight train leaving...")
        self.mediator.notify_about_departure()
