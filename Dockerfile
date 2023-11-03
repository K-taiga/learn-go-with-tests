FROM golang:1.21

WORKDIR /go/src

# godocをインストール
RUN go install golang.org/x/tools/cmd/godoc@latest

EXPOSE 8080

CMD ["sh", "-c", "GO111MODULE=off godoc -http=:8080"]