#!/bin/sh

# Wait a bit to ensure other containers are up before starting the microservices.
sleep 10

echo "Starting microservices..."

# Start each microservice
/app/pong &
/app/user &
/app/feed &
/app/publish &
/app/favorite &
/app/comment &
/app/relation &
/app/message &

echo "Microservices started."

# Wait a bit to ensure microservices are up before the API gateway.
sleep 10

echo "Starting API gateway..."

# Start the API gateway
/app/api-gateway

echo "API gateway started."