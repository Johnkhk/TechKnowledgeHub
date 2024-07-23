package flyweight

import (
	"fmt"
)

// TreeType represents the shared intrinsic state of a tree
type TreeType struct {
	Name    string
	Color   string
	Texture string
}

// Tree represents a tree object with extrinsic state (x, y) and intrinsic state (TreeType)
type Tree struct {
	X        int
	Y        int
	TreeType *TreeType
}

// Draw simulates drawing the tree on a canvas
func (t *Tree) Draw(canvas string) {
	fmt.Printf("Drawing tree on %s at (%d, %d) with type (%s, %s, %s)\n",
		canvas, t.X, t.Y, t.TreeType.Name, t.TreeType.Color, t.TreeType.Texture)
}

// TreeFactory manages the creation and reuse of TreeType instances
type TreeFactory struct {
	treeTypes []*TreeType
}

// GetTreeType returns an existing TreeType if found, or creates a new one
func (f *TreeFactory) GetTreeType(name, color, texture string) *TreeType {
	for _, treeType := range f.treeTypes {
		if treeType.Name == name && treeType.Color == color && treeType.Texture == texture {
			return treeType
		}
	}
	treeType := &TreeType{Name: name, Color: color, Texture: texture}
	f.treeTypes = append(f.treeTypes, treeType)
	return treeType
}

// Forest represents a collection of trees
type Forest struct {
	trees       []*Tree
	treeFactory *TreeFactory
}

// PlantTree plants a new tree in the forest
func (f *Forest) PlantTree(x, y int, name, color, texture string) {
	treeType := f.treeFactory.GetTreeType(name, color, texture)
	tree := &Tree{X: x, Y: y, TreeType: treeType}
	f.trees = append(f.trees, tree)
}

// Draw draws all the trees in the forest
func (f *Forest) Draw(canvas string) {
	for _, tree := range f.trees {
		tree.Draw(canvas)
	}
}
