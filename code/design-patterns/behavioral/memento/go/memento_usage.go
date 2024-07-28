package main

import "fmt"

func main() {
	document := &TextDocument{content: ""}
	editor := NewTextEditor(document)

	editor.Write("Hello, ")
	editor.Save() // Save state to memento
	editor.Write("world!")
	fmt.Println("Current State:", editor.GetContent()) // Outputs: Hello, world!

	editor.Undo()                                          // Restore previous state from memento
	fmt.Println("Restored to State:", editor.GetContent()) // Outputs: Hello,

	editor.Undo()                                          // Attempt to restore previous state, but history is empty
	fmt.Println("Restored to State:", editor.GetContent()) // Outputs: Hello, (no change)
}
