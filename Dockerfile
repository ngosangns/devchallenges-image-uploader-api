FROM golang:alpine

RUN mkdir -p /go/src/app
WORKDIR /go/src/app

RUN apk add git
RUN GO111MODULE=off go get github.com/githubnemo/CompileDaemon
RUN apk del git

ENTRYPOINT CompileDaemon -build="go build -o app" -command="./app" -polling=true -color=true