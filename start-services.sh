#!/bin/sh

# Start each microservice
/app/pong &
/app/user &
/app/feed &
/app/publish &
/app/favorite &
/app/comment &
/app/relation &
/app/message &

# Wait a bit to ensure microservices are up before the API gateway.
sleep 5

# Start the API gateway
/app/api-gateway
