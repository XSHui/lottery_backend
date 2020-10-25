#!/bin/sh
docker run  -d --name lottery_backend \
    --restart=always \
	# fix config docker daemon.json failed
    --log-driver json-file --log-opt max-size=200m \
    --net="host" \
    -v /etc/localtime:/etc/localtime:ro \
    -v /data:/data \
    xsh/lottery_backend:latest \
    --db_ip="10.23.154.125" \
    --redis_ip="10.23.101.22"
