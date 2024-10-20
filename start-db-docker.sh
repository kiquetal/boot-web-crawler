#!/bin/bash

# Set variables
CONTAINER_NAME="pg_golang"
POSTGRES_DB="chirpy"
POSTGRES_USER="wagslane"
POSTGRES_PASSWORD="paraguay"
HOST_PORT=5432

# Run PostgreSQL container
docker run --name $CONTAINER_NAME \
    -e POSTGRES_DB=$POSTGRES_DB \
    -e POSTGRES_USER=$POSTGRES_USER \
    -e POSTGRES_PASSWORD=$POSTGRES_PASSWORD \
    -p $HOST_PORT:5432 \
    -d postgres:latest

# Check if the container is running
if [ "$(docker ps -q -f name=$CONTAINER_NAME)" ]; then
    echo "PostgreSQL container is running on port $HOST_PORT"
    echo "You can connect to it using:"
    echo "Host: localhost"
    echo "Port: $HOST_PORT"
    echo "Database: $POSTGRES_DB"
    echo "Username: $POSTGRES_USER"
    echo "Password: $POSTGRES_PASSWORD"
else
    echo "Failed to start PostgreSQL container"
fi
