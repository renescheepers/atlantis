#!/bin/bash

go install

# atlantis 1
atlantis server \
  --config atlantis.yaml \
  --gh-token "$(gh auth token)" \
  --port 4141

ngrok http 4141

# atlantis 2
atlantis server \
  --config atlantis.yaml
  --gh-token "$(gh auth token)" \
  --port 4242


# redis

docker run -d --name redis-stack-server -p 6379:6379 redis/redis-stack-server:latest

# test
go run help/post_event.go
