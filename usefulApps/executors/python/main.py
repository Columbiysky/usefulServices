from typing import Union

import uvicorn
from fastapi import FastAPI, UploadFile
from fastapi.middleware.cors import CORSMiddleware
from fastapi.responses import FileResponse
from ffmpeg.videoToGif import VideoToGif
from pydantic import BaseModel


class TestObj(BaseModel):
    id: int
    name: str


app = FastAPI()

app.add_middleware(
    CORSMiddleware,
    allow_origins=["*"],
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)


@app.get('/')
def rootHello():
    return {'response': 'Hello World'}


@app.post('/')
def rootPost(obj: TestObj):
    return obj


@app.post('/videoToGif')
async def videoToGif(file: UploadFile):
    path, name, data = await VideoToGif.videoFilesInCurrentDir(file)
    return FileResponse(path=path, filename=name, media_type=data)

if __name__ == "__main__":
    uvicorn.run(app, host="0.0.0.0", port=8000)
