#! /bin/bash
docker compose build --pull --no-cache && docker compose up -d && docker ps -a