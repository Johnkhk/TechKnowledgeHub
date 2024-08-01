package main

import "fmt"

// State interface
type State interface {
	AddItem()
	SelectItem()
	InsertMoney()
	DispenseItem()
}

// VendingMachine struct
type VendingMachine struct {
	hasItemState       State
	noItemState        State
	itemRequestedState State
	hasMoneyState      State

	currentState State
}

func NewVendingMachine() *VendingMachine {
	v := &VendingMachine{}
	hasItemState := &HasItemState{vendingMachine: v}
	noItemState := &NoItemState{vendingMachine: v}
	itemRequestedState := &ItemRequestedState{vendingMachine: v}
	hasMoneyState := &HasMoneyState{vendingMachine: v}

	v.hasItemState = hasItemState
	v.noItemState = noItemState
	v.itemRequestedState = itemRequestedState
	v.hasMoneyState = hasMoneyState

	v.currentState = noItemState

	return v
}

func (v *VendingMachine) SetState(state State) {
	v.currentState = state
}

func (v *VendingMachine) AddItem() {
	v.currentState.AddItem()
}

func (v *VendingMachine) SelectItem() {
	v.currentState.SelectItem()
}

func (v *VendingMachine) InsertMoney() {
	v.currentState.InsertMoney()
}

func (v *VendingMachine) DispenseItem() {
	v.currentState.DispenseItem()
}

// Concrete States
type HasItemState struct {
	vendingMachine *VendingMachine
}

func (s *HasItemState) AddItem() {
	fmt.Println("Item already present.")
}

func (s *HasItemState) SelectItem() {
	fmt.Println("Item selected.")
	s.vendingMachine.SetState(s.vendingMachine.itemRequestedState)
}

func (s *HasItemState) InsertMoney() {
	fmt.Println("Select an item first.")
}

func (s *HasItemState) DispenseItem() {
	fmt.Println("Select an item first.")
}

type NoItemState struct {
	vendingMachine *VendingMachine
}

func (s *NoItemState) AddItem() {
	fmt.Println("Item added.")
	s.vendingMachine.SetState(s.vendingMachine.hasItemState)
}

func (s *NoItemState) SelectItem() {
	fmt.Println("No item to select.")
}

func (s *NoItemState) InsertMoney() {
	fmt.Println("No item to buy.")
}

func (s *NoItemState) DispenseItem() {
	fmt.Println("No item to dispense.")
}

type ItemRequestedState struct {
	vendingMachine *VendingMachine
}

func (s *ItemRequestedState) AddItem() {
	fmt.Println("Item already requested.")
}

func (s *ItemRequestedState) SelectItem() {
	fmt.Println("Item already selected.")
}

func (s *ItemRequestedState) InsertMoney() {
	fmt.Println("Money inserted.")
	s.vendingMachine.SetState(s.vendingMachine.hasMoneyState)
}

func (s *ItemRequestedState) DispenseItem() {
	fmt.Println("Insert money first.")
}

type HasMoneyState struct {
	vendingMachine *VendingMachine
}

func (s *HasMoneyState) AddItem() {
	fmt.Println("Item already requested.")
}

func (s *HasMoneyState) SelectItem() {
	fmt.Println("Item already selected.")
}

func (s *HasMoneyState) InsertMoney() {
	fmt.Println("Money already inserted.")
}

func (s *HasMoneyState) DispenseItem() {
	fmt.Println("Dispensing item.")
	s.vendingMachine.SetState(s.vendingMachine.noItemState)
}
