from bridge import HPPrinter, EpsonPrinter, Mac, Windows

hp_printer = HPPrinter()
epson_printer = EpsonPrinter()

mac = Mac(hp_printer)
windows = Windows(hp_printer)

mac.print()
windows.print()

mac.printer = epson_printer
windows.printer = epson_printer

mac.print()
windows.print()
