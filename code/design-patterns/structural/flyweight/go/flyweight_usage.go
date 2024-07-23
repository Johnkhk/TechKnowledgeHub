package main

func main() {
	treeFactory := &TreeFactory{}
	// pointer to Forest struct
	forest := &Forest{
		trees:       []*Tree{},
		treeFactory: treeFactory,
	}

	// These two trees share the same intrinsic state (TreeType)
	forest.PlantTree(1, 2, "Oak", "Green", "Rough")
	forest.PlantTree(3, 4, "Oak", "Green", "Rough")

	forest.PlantTree(5, 6, "Pine", "Green", "Smooth")

	canvas := "MyCanvas"
	forest.Draw(canvas)
}
