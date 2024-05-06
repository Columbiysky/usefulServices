import os

import aiofiles
from fastapi import UploadFile
from moviepy.editor import VideoFileClip


class VideoToGif:
    async def videoFilesInCurrentDir(file: UploadFile):
        dirWithInVideos = os.path.join(os.getcwd(),"InVideos")
        if not os.path.exists(dirWithInVideos):
            os.makedirs(dirWithInVideos)

        async with aiofiles.open(os.path.join(dirWithInVideos,"video.mp4"), 'wb') as out_file:
            while content := await file.read(1024):  # async read chunk
                await out_file.write(content)  # async write chunk

        videoClip = VideoFileClip(os.path.join(dirWithInVideos, "video.mp4"))
        videoClip.write_gif(os.path.join(dirWithInVideos, "videoAsGif.gif"))
        videoClip.close()

        return os.path.join(dirWithInVideos,"videoAsGif.gif"), "videoAsGif.gif", 'multipart/form-data'