#!/bin/bash 

# Stop and remove existing container if it exists
if [ "$(docker ps -q -f name=kvis-client)" ]; then
    echo "Stopping and removing existing kvis-client container..."
    docker stop kvis-client
    docker rm kvis-client
fi

# Remove existing image if it exists
if [ "$(docker images -q kvis-client)" ]; then
    echo "Removing existing kvis-client image..."
    docker rmi kvis-client
fi

# Build the new Docker image
echo "Building the new kvis-client Docker image..."
docker build -t kvis-client .

# Run the new Docker container
echo "Running the new kvis-client Docker container..."
docker run -d -p 8081:8081 --name kvis-client kvis-client