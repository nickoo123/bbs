#!/bin/bash

docker stop bbsadmins && docker rm bbsadmins
docker system prune -a
docker build -t bbsadmin .
docker run -itd --name bbsadmins -p 8080:80 --restart=always bbsadmin:latest