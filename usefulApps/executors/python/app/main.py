from typing import Union

import uvicorn
from fastapi import FastAPI
from ffmpeg.videoToGif import VideoToGif
from pydantic import BaseModel


class TestObj(BaseModel):
    id: int
    name:str

app = FastAPI()

@app.get('/')
def rootHello():
    return {'response':'Hello World'}

@app.post('/')
def rootPost(obj: TestObj):
    return obj

@app.post('/testClass')
def testClass():
    return {'response':VideoToGif.test()} 

if __name__ == "__main__":
    uvicorn.run(app, host="0.0.0.0", port=8000)