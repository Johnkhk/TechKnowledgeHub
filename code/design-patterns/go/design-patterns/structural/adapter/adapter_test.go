package adapter

import "testing"

func TestVideoPlayer(t *testing.T) {
	vp := VideoPlayer{}
	vp.Play("mp4", "example.mp4") // Should output: Playing MP4 file: example.mp4
	vp.Play("avi", "example.avi") // Should output: Playing AVI file: example.avi
}
