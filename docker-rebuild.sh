#! /bin/bash
docker stop $(docker ps -a -q) && docker rm $(docker ps -a -q) && docker system prune -a --volumes -f && docker compose build && docker compose up -d && docker ps -a