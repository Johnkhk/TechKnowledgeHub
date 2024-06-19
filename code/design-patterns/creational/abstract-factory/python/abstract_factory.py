from abc import ABC, abstractmethod


class Button(ABC):
    @abstractmethod
    def render(self):
        pass

    @abstractmethod
    def onClick(self):
        pass


class Checkbox(ABC):
    @abstractmethod
    def render(self):
        pass

    @abstractmethod
    def onChange(self):
        pass


class DarkButton(Button):
    def __init__(self, content: str):
        self.content = content

    def render(self):
        print(f"Rendering dark button with content: {self.content}")

    def onClick(self):
        print(f"Dark button with content: {self.content} was clicked")


class DarkCheckbox(Checkbox):
    def __init__(self, content: str):
        self.content = content

    def render(self):
        print(f"Rendering dark checkbox with content: {self.content}")

    def onChange(self):
        print(f"Dark checkbox with content: {self.content} was changed")


class LightButton(Button):
    def __init__(self, content: str):
        self.content = content

    def render(self):
        print(f"Rendering light button with content: {self.content}")

    def onClick(self):
        print(f"Light button with content: {self.content} was clicked")


class LightCheckbox(Checkbox):
    def __init__(self, content: str):
        self.content = content

    def render(self):
        print(f"Rendering light checkbox with content: {self.content}")

    def onChange(self):
        print(f"Light checkbox with content: {self.content} was changed")


class UIFactory(ABC):
    @abstractmethod
    def createButton(self, content: str):
        pass

    @abstractmethod
    def createCheckbox(self, content: str):
        pass


class darkUIFactory(UIFactory):
    def createButton(self, content: str):
        return DarkButton(content)

    def createCheckbox(self, content: str):
        return DarkCheckbox(content)


class lightUIFactory(UIFactory):
    def createButton(self, content: str):
        return LightButton(content)

    def createCheckbox(self, content: str):
        return LightCheckbox(content)
