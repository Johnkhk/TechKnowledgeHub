from abstract_factory import darkUIFactory, lightUIFactory


def usage(mode: str):
    dark_ui_factory = darkUIFactory()
    light_ui_factory = lightUIFactory()

    if mode == "dark":
        button = dark_ui_factory.createButton("Dark button")
        checkbox = dark_ui_factory.createCheckbox("Dark checkbox")
    elif mode == "light":
        button = light_ui_factory.createButton("Light button")
        checkbox = light_ui_factory.createCheckbox("Light checkbox")

    button.render()
    button.onClick()
    checkbox.render()
    checkbox.onChange()
