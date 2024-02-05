#! /bin/bash

docker network create mynetwork
docker-compose up --build -d
