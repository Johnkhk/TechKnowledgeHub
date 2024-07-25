from command import (
    Button,
    TurnOnCommand,
    TurnOffCommand,
    VolumeUpCommand,
    VolumeDownCommand,
)

remote_on_button = Button(TurnOnCommand())
remote_off_button = Button(TurnOffCommand())
remote_volume_up_button = Button(VolumeUpCommand())
remote_volume_down_button = Button(VolumeDownCommand())

tv_on_button = Button(TurnOnCommand())
tv_off_button = Button(TurnOffCommand())
tv_volume_up_button = Button(VolumeUpCommand())
tv_volume_down_button = Button(VolumeDownCommand())
