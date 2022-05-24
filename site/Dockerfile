FROM node:lts-alpine3.12 AS builder

MAINTAINER william

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories && \
    apk --no-cache add build-base tzdata git

COPY . /app
WORKDIR /app

RUN npm config set registry https://registry.npm.taobao.org
RUN rm -rf node_modules/ package-lock.json && yarn install && yarn build
RUN npm i -g pm2 && pm2 save

#FROM nginx:stable-alpine
#COPY --from=builder /app/dist /user/share/nginx/html
#COPY --from=builder /app/start.sh /user/share/nginx/html
#COPY --from=builder /app/.nuxt /user/share/nginx/html
#COPY --from=builder /app/static /user/share/nginx/html
#COPY --from=builder /app/nuxt.config.js /user/share/nginx/html
#COPY --from=builder /app/package.json /user/share/nginx/html

EXPOSE 3000
CMD [ "pm2-runtime", "start", "npm", "--", "start" ]
