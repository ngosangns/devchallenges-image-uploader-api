FROM golang:alpine

USER root

ENV PORT=8080

COPY . .

CMD ["./image-uploader"]