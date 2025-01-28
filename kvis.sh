#!/bin/bash

# Function to start the backend
start_backend() {
    echo "Running the backend..."
    cd backend
    nohup go run main.go > backend.log 2>&1 &
    cd ..
}

# Function to start the frontend
start_frontend() {
    echo "Running the frontend..."
    cd client
    nohup ./pipeline.sh > frontend.log 2>&1 &
    cd ..
}

# Check if the user passed "dummy" as a parameter
if [ "$1" == "dummy" ]; then
    echo "Running dummy k3d script..."
    ./dummy/run_k3d.sh
fi

# Start the backend and frontend
start_backend
start_frontend

echo "Both applications are now running."