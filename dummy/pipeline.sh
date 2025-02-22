#!/bin/bash

# filepath: /Users/joseph.baruch/REPO/kvis/dummy/run_k3d.sh

# Variables
HOSTNAME="k3d-local-registry"
IP="127.0.0.1"
HOSTS_FILE="/etc/hosts"
CLIENT_IMAGE="kvis-client:latest"
BACKEND_IMAGE="kvis-backend:latest"
REGISTRY="k3d-local-registry:5001"

# Function to start the k3d cluster
start_cluster() {
    # Check if k3d is installed
    if ! command -v k3d &> /dev/null
    then
        echo "k3d could not be found, installing k3d..."
        curl -s https://raw.githubusercontent.com/rancher/k3d/main/install.sh | bash
    else
        echo "k3d is already installed."
    fi

    # Check if the cluster is already running
    if k3d cluster list | grep -q 'dummy-cluster'
    then
        echo "k3d cluster 'dummy-cluster' is already running. Deleting the existing cluster..."
        k3d cluster delete dummy-cluster
    fi

    # Check if the local registry is already running
    if k3d registry list | grep -q 'k3d-local-registry'
    then
        echo "k3d local registry 'k3d-local-registry' is already running. Deleting the existing registry..."
        k3d registry delete k3d-local-registry
    fi

    # Create a new k3d registry
    k3d registry create local-registry --port 5001

    # Add the hostname to /etc/hosts if it doesn't exist
    if grep -q "$HOSTNAME" "$HOSTS_FILE"; then
        echo "The hostname '$HOSTNAME' already exists in $HOSTS_FILE."
    else
        echo "Adding '$HOSTNAME' to $HOSTS_FILE..."
        echo "$IP $HOSTNAME" | sudo tee -a "$HOSTS_FILE" > /dev/null

        if grep -q "$HOSTNAME" "$HOSTS_FILE"; then
            echo "Successfully added '$HOSTNAME' to $HOSTS_FILE."
        else
            echo "Failed to add '$HOSTNAME' to $HOSTS_FILE. Please check permissions."
        fi
    fi

    k3d cluster create dummy-cluster \
        --api-port 6550 \
        -p "8081:8081@loadbalancer" \
        -p "8082:8082@loadbalancer" \
        --agents 2 \
        --registry-use $REGISTRY

    # Check if the cluster is running
    if k3d cluster list | grep -q 'dummy-cluster'
    then
        echo "k3d cluster 'dummy-cluster' is up and running."
    else
        echo "Failed to create k3d cluster."
        exit 1
    fi

    echo "You can now use kubectl to interact with your k3d cluster."

    # Check if the client image exists, if not build it
    if [[ "$(docker images -q $CLIENT_IMAGE 2> /dev/null)" == "" ]]; then
        echo "Building the client Docker image..."
        cd ../client
        docker build -t $CLIENT_IMAGE .
        cd ../dummy
    fi

    # Check if the backend image exists, if not build it
    if [[ "$(docker images -q $BACKEND_IMAGE 2> /dev/null)" == "" ]]; then
        echo "Building the backend Docker image..."
        cd ../backend
        docker build -t $BACKEND_IMAGE .
        cd ../dummy
    fi

    # Tag and push the Docker images to the local registry
    docker tag $CLIENT_IMAGE $REGISTRY/$CLIENT_IMAGE
    docker tag $BACKEND_IMAGE $REGISTRY/$BACKEND_IMAGE

    docker push $REGISTRY/$CLIENT_IMAGE
    docker push $REGISTRY/$BACKEND_IMAGE

    # Apply the Kubernetes deployments and services
    kubectl apply -f ingress.yaml
    kubectl apply -f client_deployment.yaml
    kubectl apply -f client_service.yaml
    kubectl apply -f backend_deployment.yaml
    kubectl apply -f backend_service.yaml
    kubectl apply -f pvc.yaml

    # ./curl_pod/pipeline.sh
}

# Function to stop the k3d cluster
stop_cluster() {
    echo "Stopping the k3d cluster 'dummy-cluster'..."
    k3d cluster delete dummy-cluster
    echo "k3d cluster 'dummy-cluster' has been stopped."
}

# Check if the user passed "stop" as a parameter
if [ "$1" == "stop" ]; then
    stop_cluster
    exit 0
fi

# Start the k3d cluster
start_cluster