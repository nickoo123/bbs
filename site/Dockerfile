FROM node:lts-alpine3.12 AS builder

MAINTAINER william

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories && \
    apk --no-cache add build-base tzdata git

COPY . /app
WORKDIR /app

RUN npm config set registry http://registry.npmmirror.com
RUN rm -rf package-lock.json && yarn install && \
    chmod o+x /app/run.sh && \
    yarn build  && npm i -g pm2 && pm2 save
#RUN yarn add nuxt -g

EXPOSE 3000

ENTRYPOINT ["sh", "/app/run.sh"]

#FROM nginx:stable-alpine
#
#RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories && \
#    apk update && apk add nodejs && apk add npm && npm i -g pm2 && pm2 save
#COPY --from=builder /app/node_modules /user/share/nginx/html
#COPY --from=builder /app/.nuxt /user/share/nginx/html
#COPY --from=builder /app/static /user/share/nginx/html
#COPY --from=builder /app/nuxt.config.js /user/share/nginx/html
#COPY --from=builder /app/package.json /user/share/nginx/html
#
#WORKDIR /user/share/nginx/html

