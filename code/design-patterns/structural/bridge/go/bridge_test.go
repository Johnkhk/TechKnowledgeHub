package bridge

func main() {
	hpPrinter := &HPPrinter{}
	epsonPrinter := &EpsonPrinter{}

	mac := &Mac{}
	windows := &Windows{}

	mac.SetPrinter(hpPrinter)
	mac.Print()

	mac.SetPrinter(epsonPrinter)
	mac.Print()

	windows.SetPrinter(hpPrinter)
	windows.Print()

	windows.SetPrinter(epsonPrinter)
	windows.Print()
}
