FROM golang:default

WORKDIR /go/src

RUN go env -w GO111MODULE="on"
RUN go env -w GOPROXY="https://goproxy.cn,direct"
RUN go mod init myserver

RUN go get -u -v github.com/gin-gonic/gin
RUN go get -u -v go.mongodb.org/mongo-driver/mongo

