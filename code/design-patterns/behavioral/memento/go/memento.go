package main

// TextDocument represents the document being edited
type TextDocument struct {
	content string
}

func (d *TextDocument) GetContent() string {
	return d.content
}

func (d *TextDocument) SetContent(content string) {
	d.content = content
}

// Memento encapsulates the state of the TextDocument
type Memento struct {
	state string
}

func NewMemento(state string) *Memento {
	return &Memento{state: state}
}

func (m *Memento) GetState() string {
	return m.state
}

// TextEditor acts as both the Originator and the Caretaker
type TextEditor struct {
	document *TextDocument
	history  []*Memento
}

func NewTextEditor(document *TextDocument) *TextEditor {
	return &TextEditor{
		document: document,
		history:  []*Memento{},
	}
}

func (e *TextEditor) Write(text string) {
	e.document.SetContent(e.document.GetContent() + text)
}

func (e *TextEditor) Save() {
	e.history = append(e.history, NewMemento(e.document.GetContent()))
}

func (e *TextEditor) Undo() {
	if len(e.history) > 0 {
		memento := e.history[len(e.history)-1]
		e.history = e.history[:len(e.history)-1]
		e.document.SetContent(memento.GetState())
	}
}

func (e *TextEditor) GetContent() string {
	return e.document.GetContent()
}
