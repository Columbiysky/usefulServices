import subprocess
import os
from videoFormatsEnum import VideoFormat

def videoFilesInCurrentDir():
    files = (os.listdir(os.getcwd()))
    videoFilesEnum = [i.name for i in VideoFormat]
    result = []
    for file in files:
        if file.split('.')[-1] in videoFilesEnum:
            result.append(file)
    return result


if __name__ == '__main__':
    videoFiles = videoFilesInCurrentDir()
    for file in videoFiles:
        command = 'ffmpeg -i ' + file +' '+ file + '.gif'
        subprocess.call(command, shell=True)