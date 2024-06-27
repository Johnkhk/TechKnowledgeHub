package bridge

import "fmt"

// Abstraction Layer

type Computer struct {
	printer Printer
}

func (c *Computer) SetPrinter(printer Printer) {
	c.printer = printer
}

func (c *Computer) Print() {
	c.printer.Print()
}

type Mac struct {
	Computer
}

type Windows struct {
	Computer
}

// Implementor Layer

type Printer interface {
	Print()
}

type HPPrinter struct{}

func (p *HPPrinter) Print() {
	fmt.Println("Printing by an HP printer")
}

type EpsonPrinter struct{}

func (p *EpsonPrinter) Print() {
	fmt.Println("Printing by an Epson printer")
}
