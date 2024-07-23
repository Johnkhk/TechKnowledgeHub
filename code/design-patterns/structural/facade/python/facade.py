class TV:
    def turn_on(self):
        print("TV is ON")

    def turnOff(self):
        print("TV is OFF")


class SoundSystem:
    def turn_on(self):
        print("Sound System is ON")

    def turnOff(self):
        print("Sound System is OFF")


class GamingConsole:
    def turn_on(self):
        print("Gaming Console is ON")

    def turnOff(self):
        print("Gaming Console is OFF")


class HomeEntertainmentFacade:

    def __init__(self):
        self.tv = TV()
        self.sound_system = SoundSystem()
        self.gaming_console = GamingConsole()

    def playGame(self):
        self.gaming_console.turn_on()
        self.tv.turn_on()
        self.sound_system.turn_on()

    def watchMovie(self):
        self.tv.turn_on()
        self.sound_system.turn_on()
