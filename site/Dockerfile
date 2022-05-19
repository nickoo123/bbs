FROM node:apline AS builder

MAINTAINER william

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
RUN apk --no-cache add build-base tzdata git

COPY . /app
WORKDIR /app

RUN npm config set registry https://registry.npm.taobao.org && \
    yarn install
