FROM golang:1.22.1-alpine3.19
WORKDIR /

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY *.go ./
COPY ./controllers/account/*.go ./controllers/account/
COPY ./controllers/token/*.go ./controllers/token/

COPY ./dbLogic/account/*.go ./dbLogic/account/
COPY ./dbLogic/token/*.go ./dbLogic/token/
COPY ./dbLogic/base/*.go ./dbLogic/base/

COPY ./logic/account/*.go ./logic/account/
COPY ./logic/token/*.go ./logic/token/

COPY ./models/*.go ./models/
COPY ./tests/*.go ./tests/

RUN go build -v .

EXPOSE 30001