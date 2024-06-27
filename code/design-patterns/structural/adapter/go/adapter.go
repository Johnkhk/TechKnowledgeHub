package adapter

import "fmt"

// MP4Player is a simple struct to play MP4 files.
type MP4Player struct{}

// PlayMP4 is a method to play MP4 files.
func (mp *MP4Player) PlayMP4(fileName string) {
	fmt.Printf("Playing MP4 file: %s\n", fileName)
}

// AVIPlayer is a simple struct to play AVI files.
type AVIPlayer struct{}

// PlayAVI is a method to play AVI files.
func (av *AVIPlayer) PlayAVI(fileName string) {
	fmt.Printf("Playing AVI file: %s\n", fileName)
}

// MediaAdapter is an adapter for media players that do not match the required method signature.
type MediaAdapter struct {
	player *AVIPlayer
}

// NewMediaAdapter creates a new MediaAdapter instance based on the fileType.
func NewMediaAdapter(fileType string) *MediaAdapter {
	if fileType == "avi" {
		return &MediaAdapter{player: &AVIPlayer{}}
	}
	return nil
}

// Play checks the fileType and plays the file using the appropriate method.
func (ma *MediaAdapter) Play(fileType, fileName string) {
	if fileType == "avi" && ma.player != nil {
		ma.player.PlayAVI(fileName)
	}
}

// VideoPlayer can play different formats by using the appropriate player or adapter.
type VideoPlayer struct{}

// Play determines how to play the media based on the file type.
func (vp *VideoPlayer) Play(fileType, fileName string) {
	if fileType == "mp4" {
		player := MP4Player{}
		player.PlayMP4(fileName)
	} else {
		adapter := NewMediaAdapter(fileType)
		if adapter != nil {
			adapter.Play(fileType, fileName)
		} else {
			fmt.Println("Unsupported file type")
		}
	}
}
