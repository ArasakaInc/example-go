FROM docker.zhubby.com/library/golang:1.21 as build
WORKDIR /go/src/
ADD . .
RUN  unset GOPATH && go build -v -mod=vendor -o main
ENTRYPOINT ["./main"]
