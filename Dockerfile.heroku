FROM golang:alpine

RUN mkdir /go/src/app
WORKDIR /go/src/app

COPY . .

CMD go build -o app; ./app