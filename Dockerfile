FROM golang:alpine

USER root
ENV PORT=8080

ADD . /go/src/app
WORKDIR /go/src/app

RUN go build -o image-uploader .

CMD /go/src/app/image-uploader