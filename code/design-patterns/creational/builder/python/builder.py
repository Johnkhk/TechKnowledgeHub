from abc import ABC, abstractmethod


class House:
    def __init__(self):
        self.foundation = None
        self.structure = None
        self.roof = None
        self.rooms = None
        self.backyard = None

    def __str__(self):
        return (
            f"House with foundation: {self.foundation}, structure: {self.structure}, "
            f"roof: {self.roof}, rooms: {self.rooms}, backyard: {self.backyard}"
        )


class HouseBuilder(ABC):
    @abstractmethod
    def build_foundation(self):
        pass

    @abstractmethod
    def build_structure(self):
        pass

    @abstractmethod
    def build_roof(self):
        pass

    @abstractmethod
    def build_rooms(self, number):
        pass

    @abstractmethod
    def build_backyard(self, type):
        pass

    @abstractmethod
    def build(self):
        pass


class WoodenHouseBuilder(HouseBuilder):
    def __init__(self):
        self.house = House()

    def build_foundation(self):
        self.house.foundation = "Wooden Foundation"
        return self

    def build_structure(self):
        self.house.structure = "Wooden Structure"
        return self

    def build_roof(self):
        self.house.roof = "Wooden Roof"
        return self

    def build_rooms(self, number):
        self.house.rooms = f"{number} Wooden Rooms"
        return self

    def build_backyard(self, type):
        self.house.backyard = f"Wooden {type} Backyard"
        return self

    def build(self):
        return self.house


class BrickHouseBuilder(HouseBuilder):
    def __init__(self):
        self.house = House()

    def build_foundation(self):
        self.house.foundation = "Concrete Foundation"
        return self

    def build_structure(self):
        self.house.structure = "Brick Structure"
        return self

    def build_roof(self):
        self.house.roof = "Concrete Roof"
        return self

    def build_rooms(self, number):
        self.house.rooms = f"{number} Brick Rooms"
        return self

    def build_backyard(self, type):
        self.house.backyard = f"Brick {type} Backyard"
        return self

    def build(self):
        return self.house


class Director:
    def __init__(self, builder):
        self._builder = builder

    def construct_simple_house(self):
        self._builder.build_foundation().build_structure().build_roof().build_rooms(2)
        return self._builder.build()

    def construct_luxury_house(self):
        self._builder.build_foundation().build_structure().build_roof().build_rooms(
            5
        ).build_backyard("Luxury")
        return self._builder.build()
