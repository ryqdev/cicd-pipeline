#!/bin/bash

# Exit immediately if a command exits with a non-zero status
set -e

# Define variables
IMAGE_NAME="cicd-pipeline"
CONTAINER_NAME="cicd-pipeline-container"
PORT=8080

echo "Building Docker image..."
docker build -t $IMAGE_NAME .

# Check if a container with the same name is already running
if docker ps -a | grep -q $CONTAINER_NAME; then
    echo "Stopping and removing existing container..."
    docker stop $CONTAINER_NAME || true
    docker rm $CONTAINER_NAME || true
fi

echo "Starting container..."
docker run -d \
    --name $CONTAINER_NAME \
    -p $PORT:$PORT \
    $IMAGE_NAME

echo "Container started successfully!"
echo "You can access the application at: http://localhost:$PORT"
echo "Available endpoints:"
echo "  - GET  http://localhost:$PORT/"
echo "  - GET  http://localhost:$PORT/ping"
echo "  - POST http://localhost:$PORT/webhook"