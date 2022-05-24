#!/bin/bash

set -eu

docker build -t bbs .

docker run -itd --name bbsgo -p 8082:8082 --restart=always -v /data/bbs:/data bbs:latest