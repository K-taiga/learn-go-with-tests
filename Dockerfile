FROM golang:1.21

WORKDIR /go/src/app

COPY . .

EXPOSE 8080