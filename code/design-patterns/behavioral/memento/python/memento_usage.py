from memento import TextDocument, TextEditor


document = TextDocument("")
editor = TextEditor(document)

editor.write("Hello, ")
editor.save()  # Save state to memento
editor.write("world!")
print("Current State:", editor.get_content())  # Outputs: Hello, world!

editor.undo()  # Restore previous state from memento
print("Restored to State:", editor.get_content())  # Outputs: Hello,

editor.undo()  # Attempt to restore previous state, but history is empty
print("Restored to State:", editor.get_content())  # Outputs: Hello, (no change)
