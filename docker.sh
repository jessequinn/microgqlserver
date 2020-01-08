#!/bin/sh
docker-compose stop
docker-compose up --build --remove-orphans -d
sleep 2
#docker-compose exec mongo1 mongo /root/000_init_replicaSet.js
docker-compose logs -f srv cli api
