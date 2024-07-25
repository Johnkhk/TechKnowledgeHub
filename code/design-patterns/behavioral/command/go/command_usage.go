package main

func main() {
	turnOnCommand := &TurnOnCommand{}
	turnOffCommand := &TurnOffCommand{}
	volumeUpCommand := &VolumeUpCommand{}
	volumeDownCommand := &VolumeDownCommand{}

	onButton := NewButton(turnOnCommand)
	offButton := NewButton(turnOffCommand)
	volumeUpButton := NewButton(volumeUpCommand)
	volumeDownButton := NewButton(volumeDownCommand)

	onButton.OnClick()
	offButton.OnClick()
	volumeUpButton.OnClick()
	volumeDownButton.OnClick()
}
