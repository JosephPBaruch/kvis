#!/bin/bash 

IMAGE="kvis-client"

# Stop and remove existing container if it exists
if [ "$(docker ps -q -f name=$IMAGE)" ]; then
    echo "Stopping and removing existing $IMAGE container..."
    docker stop $IMAGE
    docker rm $IMAGE
fi

# Remove existing image if it exists
if [ "$(docker images -q $IMAGE)" ]; then
    echo "Removing existing $IMAGE image..."
    docker rmi $IMAGE
fi

# Build the new Docker image
echo "Building the new $IMAGE Docker image..."
docker build -t $IMAGE .

# Run the new Docker container
echo "Running the new $IMAGE Docker container..."
docker run -d -p 8081:8081 --name $IMAGE $IMAGE