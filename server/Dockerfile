FROM golang:1.17-alpine AS builder

MAINTAINER william

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories && \
    apk --no-cache add build-base tzdata git

COPY . /app
WORKDIR /app
RUN go env -w GOPROXY=https://goproxy.cn,direct && \
    go env -w GO111MODULE=on && \
    go mod tidy && go mod vendor && \
    GOOS=linux GOARCH=amd64 go build

RUN pwd && ls

FROM alpine

RUN mkdir -p /data && mkdir -p /data/www && mkdir -p /data/logs
WORKDIR /app
COPY --from=builder /app/bbs /app
COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /etc/localtime

EXPOSE 8082

CMD ["./bbs"]