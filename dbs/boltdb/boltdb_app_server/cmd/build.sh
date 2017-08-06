#!/bin/bash
GOOS=linux go build -o mybin *go
docker build -t hippoai/later-boltdb-app-server:latest . 
