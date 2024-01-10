cd usefulApps

cd backend/src
go get .

cd ../..
cd executors/python
pip install -r ./requirements.txt -U

cd ../..
cd frontend
npm i

cd ../..
cd sso
go get .

cd ../..
cd tools/dbTool
go get .