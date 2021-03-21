# 基本的依赖下载
FROM golang:1.14.2-alpine3.11 AS builder
RUN apk update && apk upgrade && apk add --no-cache bash git openssh
RUN go get -v -u github.com/nsqio/go-nsq

# 拷贝源代码并编译
FROM builder AS compiler
RUN mkdir -p /go/src/app
WORKDIR /go/src/app
COPY . .
RUN go build .

# 生成可执行的镜像
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=compiler /go/src/app/app .

ENTRYPOINT ["./app"]