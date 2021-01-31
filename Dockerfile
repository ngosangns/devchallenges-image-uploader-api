FROM golang:alpine

USER root

ENV PORT=8080

COPY . .

RUN go build

CMD ./image-uploader