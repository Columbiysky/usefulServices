cd ..\..\usefulApps
cd backend\src
go get .
go mod tidy

cd ..\..
cd sso
go get .
go mod tidy

cd ..
cd services\python
pip install -r .\requirements.txt -U

cd ..\..
cd frontend
npm i

cd ..\..
cd services\ts
npm i

cd ..\..
cd sso
go get .
go mod tidy

cd ..\..
cd tools\dbTool
go get .
go mod tidy

echo "finished..."
pause