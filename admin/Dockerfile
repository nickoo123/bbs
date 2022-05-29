FROM node:12.3.1-alpine AS builder

MAINTAINER william

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories && \
    apk --no-cache add build-base tzdata git

COPY . /app
WORKDIR /app

RUN npm config set registry https://registry.npm.taobao.org -g && apk add python2 && \
    rm -rf node_modules && yarn install && yarn build

FROM nginx:stable-alpine
COPY --from=builder /app/dist /usr/share/nginx/html
WORKDIR /usr/share/nginx/html
