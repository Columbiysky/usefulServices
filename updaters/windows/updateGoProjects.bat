cd ..\..\usefulApps
cd backend\src
go get .
go mod tidy

cd ..\..
cd sso
go get .
go mod tidy

pause