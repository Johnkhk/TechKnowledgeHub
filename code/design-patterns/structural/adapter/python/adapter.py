class MP4Player:
    def playMP4(self, fileName):
        print(f"Playing MP4 file: {fileName}")


class AVIPlayer:
    def playAVI(self, fileName):
        print(f"Playing AVI file: {fileName}")


class MediaAdapter:
    def __init__(self, fileType):
        if fileType == "avi":
            self.player = AVIPlayer()
        else:
            self.player = None

    def play(self, fileType, fileName):
        if fileType == "mp4":
            MP4Player().playMP4(fileName)
        elif fileType == "avi":
            self.player.playAVI(fileName)


class VideoPlayer:
    def play(self, fileType, fileName):
        if fileType == "mp4":
            MP4Player().playMP4(fileName)
        else:
            adapter = MediaAdapter(fileType)
            adapter.play(fileType, fileName)
