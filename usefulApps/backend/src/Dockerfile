FROM golang:1.22.1-alpine3.19
WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY *.go ./
COPY ./controllers/account/*.go ./controllers/account/
COPY ./controllers/*.go ./controllers/

COPY ./dbLogic/account/*.go ./dbLogic/account/
COPY ./dbLogic/base/*.go ./dbLogic/base/

COPY ./logic/account/*.go ./logic/account/
COPY ./logic/hello/params/*.go ./logic/hello/params/
COPY ./logic/hello/*.go ./logic/hello/

COPY ./models/*.go ./models/

RUN go build -v .

EXPOSE 30000