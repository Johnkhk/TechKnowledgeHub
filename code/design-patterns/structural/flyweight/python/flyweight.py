from typing import List


class TreeType:
    def __init__(self, name, color, texture):
        self.name = name
        self.color = color
        self.texture = texture


class Tree:
    def __init__(self, x, y, tree_type: TreeType):
        self.x = x
        self.y = y
        self.tree_type = tree_type

    def draw(self, canvas, x, y):
        print("Drawing on canvas")


class TreeFactory:
    tree_types: List[TreeType] = []

    def get_tree_type(self, name, color, texture) -> TreeType:
        # Look for existing tree type
        for tree_type in self.tree_types:
            if (
                tree_type.name == name
                and tree_type.color == color
                and tree_type.texture == texture
            ):
                return tree_type
        # If not found, create a new one and add it to the list
        tree_type = TreeType(name, color, texture)
        self.tree_types.append(tree_type)
        return tree_type


class Forest:

    def __init__(self):
        self.trees: List[Tree] = []

    def plant_tree(self, x, y, name, color, texture) -> Tree:
        tree_type = TreeFactory().get_tree_type(name, color, texture)
        tree = Tree(x, y, tree_type)
        self.trees.append(tree)
        return tree

    def draw(self, canvas):
        for tree in self.trees:
            tree.draw(canvas, tree.x, tree.y)
