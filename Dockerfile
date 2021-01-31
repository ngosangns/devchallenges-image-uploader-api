FROM golang:alpine
USER root

# Move files to Go source path
ADD . /go/src/app
WORKDIR /go/src/app

# Build and run app
RUN go build -o app .
CMD ["/go/src/app/app"]