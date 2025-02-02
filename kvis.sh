#!/bin/bash

# Function to start the backend
start_backend() {
    echo "Running the backend..."
    cd backend
    nohup go run main.go > backend.log 2>&1 &
    BACKEND_PID=$!
    cd ..
}

# Function to start the frontend
start_frontend() {
    echo "Running the frontend..."
    cd client
    npm run dev &
    FRONTEND_PID=$!
    cd ..
}

# Function to clean up processes
cleanup() {
    echo "Stopping the backend and frontend..."
    kill $BACKEND_PID $FRONTEND_PID
    echo "Stopping the cluster..."
    cd dummy
    ./pipeline.sh stop
    cd ..
    echo "Cleanup complete."
    exit 0
}

# Trap SIGINT and SIGTERM to run cleanup
trap cleanup SIGINT SIGTERM

# Check if the user passed "dummy" as a parameter
if [ "$1" == "dummy" ]; then
    echo "Running dummy k3d script..."
    cd dummy
    ./pipeline.sh
    cd ..
fi

# Start the backend and frontend
start_backend
start_frontend

echo "Both applications are now running."
echo "Visit http://localhost:8081 to see the frontend."

# Wait indefinitely to keep the script running
while true; do
    sleep 1
done