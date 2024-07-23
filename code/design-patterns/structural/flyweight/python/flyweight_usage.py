from flyweight import Forest

forest = Forest()
# The same tree type is shared between the two trees
forest.plant_tree(0, 0, "Oak", "Green", "Rough")
forest.plant_tree(1, 1, "Oak", "Green", "Rough")
