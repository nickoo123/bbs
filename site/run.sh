#!/bin/sh

pm2 start npm --name 'eight_rice' -- run start &

while true
do
  sleep 30
done
