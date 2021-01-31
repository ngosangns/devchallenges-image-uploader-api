FROM golang:alpine

ENV PORT=8080

COPY . .

CMD ["./image-uploader"]