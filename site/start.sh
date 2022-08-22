#!/bin/bash

set -eu

docker system prune -a
docker build -t bbssite .
docker stop bbssites && docker rm bbssites
docker run -itd --name bbssites -p 3000:3000 --restart=always bbssite:latest
