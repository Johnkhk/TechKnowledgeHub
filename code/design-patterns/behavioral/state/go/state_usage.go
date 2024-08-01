package main

// Example usage
func main() {
	vendingMachine := NewVendingMachine()

	vendingMachine.AddItem()
	vendingMachine.SelectItem()
	vendingMachine.InsertMoney()
	vendingMachine.DispenseItem()
	vendingMachine.AddItem()
}
