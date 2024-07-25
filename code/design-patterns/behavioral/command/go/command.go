package main

import (
	"fmt"
)

// Command interface
type Command interface {
	Execute()
}

// Button struct
type Button struct {
	command Command
}

// NewButton creates a new Button with a given command
func NewButton(command Command) *Button {
	return &Button{command: command}
}

// OnClick calls the execute method of the command
func (b *Button) OnClick() {
	b.command.Execute()
}

// TurnOnCommand struct
type TurnOnCommand struct{}

// Execute method for TurnOnCommand
func (c *TurnOnCommand) Execute() {
	fmt.Println("Turning on TV")
}

// TurnOffCommand struct
type TurnOffCommand struct{}

// Execute method for TurnOffCommand
func (c *TurnOffCommand) Execute() {
	fmt.Println("Turning off TV")
}

// VolumeUpCommand struct
type VolumeUpCommand struct{}

// Execute method for VolumeUpCommand
func (c *VolumeUpCommand) Execute() {
	fmt.Println("Turning up volume")
}

// VolumeDownCommand struct
type VolumeDownCommand struct{}

// Execute method for VolumeDownCommand
func (c *VolumeDownCommand) Execute() {
	fmt.Println("Turning down volume")
}
