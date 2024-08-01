from state import VendingMachine

# Example usage
vendingMachine = VendingMachine()

vendingMachine.addItem()
print(vendingMachine.currentState)
vendingMachine.selectItem()
print(vendingMachine.currentState)
vendingMachine.insertMoney()
print(vendingMachine.currentState)
vendingMachine.dispenseItem()
print(vendingMachine.currentState)
vendingMachine.addItem()
