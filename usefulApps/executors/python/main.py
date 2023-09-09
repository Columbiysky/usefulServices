from typing import Union

import uvicorn
from fastapi import FastAPI
from fastapi.middleware.cors import CORSMiddleware
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


@app.post('/testClass')
def testClass():
    return {'response': VideoToGif.test()}


if __name__ == "__main__":
    uvicorn.run(app, host="0.0.0.0", port=8000)
