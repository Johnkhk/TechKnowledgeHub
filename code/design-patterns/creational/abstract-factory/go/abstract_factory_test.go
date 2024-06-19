package abstract_factory

func main() {
	var factory UIFactory

	// Use the DarkUIFactory
	factory = DarkUIFactory{}
	button := factory.CreateButton("Dark Mode")
	checkbox := factory.CreateCheckbox("Dark Mode")
	button.Render()
	button.OnClick()
	checkbox.Render()
	checkbox.OnChange()

	// Use the LightUIFactory
	factory = LightUIFactory{}
	button = factory.CreateButton("Light Mode")
	checkbox = factory.CreateCheckbox("Light Mode")
	button.Render()
	button.OnClick()
	checkbox.Render()
	checkbox.OnChange()
}
