#!/bin/bash

IMAGE_NAME="kvis-backend"
CONTAINER_NAME="kvis-backend"
DOCKERFILE_PATH="."

# Check if the container is already running
if [[ "$(docker ps -q -f name=$CONTAINER_NAME)" ]]; then
    echo "Stopping the running container..."
    docker stop $CONTAINER_NAME
    echo "Removing the container..."
    docker rm $CONTAINER_NAME
fi

# Check if the image exists
if [[ "$(docker images -q $IMAGE_NAME 2> /dev/null)" != "" ]]; then
    echo "Removing the existing Docker image..."
    docker rmi $IMAGE_NAME
fi

# Build the Docker image
echo "Building the Docker image..."
docker build -t $IMAGE_NAME $DOCKERFILE_PATH

# Run the container
echo "Starting the container..."
docker run -d -p 8082:8082 --name $CONTAINER_NAME $IMAGE_NAME

echo "Container is running and accessible at http://localhost:8082"