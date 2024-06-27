from abc import ABC, abstractmethod


class Printer(ABC):
    @abstractmethod
    def print(self):
        pass


class HPPrinter(Printer):
    def print(self):
        print("Printing by an HP printer")


class EpsonPrinter(Printer):
    def print(self):
        print("Printing by an Epson printer")


class Computer(ABC):
    def __init__(self, printer: Printer):
        self.printer = printer

    @abstractmethod
    def print(self):
        pass


class Mac(Computer):
    def print(self):
        print("Mac computer printing using", end=" ")
        self.printer.print()


class Windows(Computer):
    def print(self):
        print("Windows computer printing using", end=" ")
        self.printer.print()
