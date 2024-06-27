from adapter import VideoPlayer

# Create an instance of VideoPlayer
video_player = VideoPlayer()

# Play an MP4 file
print("Testing MP4 Playback:")
video_player.play("mp4", "sample.mp4")

# Play an AVI file
print("\nTesting AVI Playback:")
video_player.play("avi", "example.avi")
