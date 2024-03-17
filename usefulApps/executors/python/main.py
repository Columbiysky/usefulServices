import json
from typing import Union

import requests
import uvicorn
from fastapi import FastAPI, Request, UploadFile
from fastapi.middleware.cors import CORSMiddleware
from fastapi.responses import FileResponse, JSONResponse
from pydantic import BaseModel
from videoToGif.videoToGif import VideoToGif


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
    return JSONResponse({'response': 'Hello World'})

@app.post('/')
def rootPost(obj: TestObj):
    return obj

@app.post('/videoToGif')
async def videoToGif(file: UploadFile):
    path, name, data = await VideoToGif.videoFilesInCurrentDir(file)
    return FileResponse(path=path, filename=name, media_type=data)

@app.middleware("http")
async def auth_middleware(request: Request, call_next):
    tokenArr = request.headers.get('Authorization').split(' ')
    while('' in tokenArr):
        tokenArr.remove('')

    if len(tokenArr) < 2:
        return JSONResponse('Unauthorized', 401)
    
    token = tokenArr[1]

    if token != 'CI_TEST':
        headers = {'Authorization' : 'Bearer ' + token}
        ssoReq = requests.get('http://127.0.0.1:8081/checkToken', headers = headers)
        tokenStatus = json.loads(ssoReq.text)["status"]

        if tokenStatus != 'exists' :
            return JSONResponse('Unauthorized', 401)

    response = await call_next(request)
    return response

if __name__ == "__main__":
    uvicorn.run(app, host="0.0.0.0", port=8000)
