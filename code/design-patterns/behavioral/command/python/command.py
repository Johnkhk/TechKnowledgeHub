from abc import ABC, abstractmethod


class Command(ABC):
    @abstractmethod
    def execute(self):
        pass


class Button:
    def __init__(self, command: Command):
        self.command = command

    def onClick(self):
        self.command.execute()


class TurnOnCommand(Command):
    def execute(self):
        print("Turning on TV")


class TurnOffCommand(Command):
    def execute(self):
        print("Turning off TV")


class VolumeUpCommand(Command):
    def execute(self):
        print("Turning up volume")


class VolumeDownCommand(Command):
    def execute(self):
        print("Turning down volume")
