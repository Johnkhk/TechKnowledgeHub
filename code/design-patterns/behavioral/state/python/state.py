from abc import ABC, abstractmethod


# State Interface
class State(ABC):
    def __init__(self, vendingMachine):
        self.vendingMachine = vendingMachine

    @abstractmethod
    def addItem(self):
        pass

    @abstractmethod
    def selectItem(self):
        pass

    @abstractmethod
    def insertMoney(self):
        pass

    @abstractmethod
    def dispenseItem(self):
        pass


# Concrete States
class HasItem(State):
    def addItem(self):
        print("Item already present.")

    def selectItem(self):
        print("Item selected.")
        self.vendingMachine.setState(self.vendingMachine.itemRequestedState)

    def insertMoney(self):
        print("Select an item first.")

    def dispenseItem(self):
        print("Select an item first.")


class NoItem(State):
    def addItem(self):
        print("Item added.")
        self.vendingMachine.setState(self.vendingMachine.hasItemState)

    def selectItem(self):
        print("No item to select.")

    def insertMoney(self):
        print("No item to buy.")

    def dispenseItem(self):
        print("No item to dispense.")


class ItemRequested(State):
    def addItem(self):
        print("Item already requested.")

    def selectItem(self):
        print("Item already selected.")

    def insertMoney(self):
        print("Money inserted.")
        self.vendingMachine.setState(self.vendingMachine.hasMoneyState)

    def dispenseItem(self):
        print("Insert money first.")


class HasMoney(State):
    def addItem(self):
        print("Item already requested.")

    def selectItem(self):
        print("Item already selected.")

    def insertMoney(self):
        print("Money already inserted.")

    def dispenseItem(self):
        print("Dispensing item.")
        self.vendingMachine.setState(self.vendingMachine.noItemState)


# Vending Machine as a Context and a State
class VendingMachine:
    def __init__(self):
        self.hasItemState = HasItem(self)
        self.noItemState = NoItem(self)
        self.itemRequestedState = ItemRequested(self)
        self.hasMoneyState = HasMoney(self)

        self.currentState = self.noItemState

    def setState(self, state):
        self.currentState = state

    def addItem(self):
        self.currentState.addItem()

    def selectItem(self):
        self.currentState.selectItem()

    def insertMoney(self):
        self.currentState.insertMoney()

    def dispenseItem(self):
        self.currentState.dispenseItem()
