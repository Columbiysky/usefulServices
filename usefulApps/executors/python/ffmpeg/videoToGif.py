import os

from ffmpeg.videoFormatsEnum import VideoFormat


class VideoToGif:
    def videoFilesInCurrentDir():
        files = (os.listdir(os.getcwd()))
        videoFilesEnum = [i.name for i in VideoFormat]
        result = []
        for file in files:
            if file.split('.')[-1] in videoFilesEnum:
                result.append(file)
        return result
    
    def test ():
        return 'test VideoToGif'