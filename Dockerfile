FROM golang:alpine

RUN apk add git; GO111MODULE=off go get github.com/githubnemo/CompileDaemon

RUN mkdir /go/src/app
WORKDIR /go/src/app

ENTRYPOINT CompileDaemon -build="go build -o app" -command="./app" -polling=true