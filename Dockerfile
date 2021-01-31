FROM golang:alpine

USER root

ENV PORT=8080

COPY . .

RUN go build -o image-uploader .

CMD ./image-uploader