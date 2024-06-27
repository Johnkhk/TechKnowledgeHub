package abstract_factory

import "fmt"

// Button interface
type Button interface {
	Render()
	OnClick()
}

// Checkbox interface
type Checkbox interface {
	Render()
	OnChange()
}

// DarkButton struct
type DarkButton struct {
	Content string
}

func (b DarkButton) Render() {
	fmt.Printf("Rendering dark button with content: %s\n", b.Content)
}

func (b DarkButton) OnClick() {
	fmt.Printf("Dark button with content: %s was clicked\n", b.Content)
}

// DarkCheckbox struct
type DarkCheckbox struct {
	Content string
}

func (c DarkCheckbox) Render() {
	fmt.Printf("Rendering dark checkbox with content: %s\n", c.Content)
}

func (c DarkCheckbox) OnChange() {
	fmt.Printf("Dark checkbox with content: %s was changed\n", c.Content)
}

// LightButton struct
type LightButton struct {
	Content string
}

func (b LightButton) Render() {
	fmt.Printf("Rendering light button with content: %s\n", b.Content)
}

func (b LightButton) OnClick() {
	fmt.Printf("Light button with content: %s was clicked\n", b.Content)
}

// LightCheckbox struct
type LightCheckbox struct {
	Content string
}

func (c LightCheckbox) Render() {
	fmt.Printf("Rendering light checkbox with content: %s\n", c.Content)
}

func (c LightCheckbox) OnChange() {
	fmt.Printf("Light checkbox with content: %s was changed\n", c.Content)
}

// UIFactory interface
type UIFactory interface {
	CreateButton(content string) Button
	CreateCheckbox(content string) Checkbox
}

// DarkUIFactory struct
type DarkUIFactory struct{}

func (f DarkUIFactory) CreateButton(content string) Button {
	return DarkButton{Content: content}
}

func (f DarkUIFactory) CreateCheckbox(content string) Checkbox {
	return DarkCheckbox{Content: content}
}

// LightUIFactory struct
type LightUIFactory struct{}

func (f LightUIFactory) CreateButton(content string) Button {
	return LightButton{Content: content}
}

func (f LightUIFactory) CreateCheckbox(content string) Checkbox {
	return LightCheckbox{Content: content}
}
