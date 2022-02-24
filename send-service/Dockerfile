# 发送服务
FROM golang:1.17-alpine

WORKDIR /go/src/app
COPY . .
ENV GOPROXY="https://goproxy.cn,direct"

RUN go mod download
RUN go build -o main server.go

EXPOSE 9002

CMD ["./main"]