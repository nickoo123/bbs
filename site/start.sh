#!/bin/bash

docker stop bbssites && docker rm bbssites
docker system prune -a
docker build -t bbssite .
docker run -itd --name bbsites -p 3000:3000 --restart=always bbssite:latest
