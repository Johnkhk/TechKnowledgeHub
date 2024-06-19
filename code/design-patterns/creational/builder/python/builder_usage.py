from builder import WoodenHouseBuilder, BrickHouseBuilder, Director

# Client code
wooden_house_builder = WoodenHouseBuilder()
brick_house_builder = BrickHouseBuilder()

director = Director(wooden_house_builder)
simple_wooden_house = director.construct_simple_house()
print(simple_wooden_house)

director = Director(brick_house_builder)
luxury_brick_house = director.construct_luxury_house()
print(luxury_brick_house)
