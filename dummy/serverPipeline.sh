#!/bin/bash

# filepath: /Users/joseph.baruch/REPO/kvis/dummy/run_k3d.sh

# Variables
HOSTNAME="localhost:5000"
IP="127.0.0.1"
HOSTS_FILE="/etc/hosts"
CLIENT_IMAGE="kvis-client:latest"
BACKEND_IMAGE="kvis-backend:latest"
REGISTRY="localhost:5000"

cleanup() {
    echo "Cleaning up existing resources..."

    # Delete the registry if it exists
    if [ "$(docker ps -q -f name=registry)" ]; then
        docker stop registry
        docker rm registry
    fi

    # Delete the images if they exist
    if [ "$(docker images -q $CLIENT_IMAGE 2> /dev/null)" ]; then
        docker rmi $CLIENT_IMAGE
    fi

    # if [ "$(docker images -q $BACKEND_IMAGE 2> /dev/null)" ]; then
    #     docker rmi $BACKEND_IMAGE
    # fi

    if [ "$(docker images -q $REGISTRY/$CLIENT_IMAGE 2> /dev/null)" ]; then
        docker rmi $REGISTRY/$CLIENT_IMAGE
    fi

    # if [ "$(docker images -q $REGISTRY/$BACKEND_IMAGE 2> /dev/null)" ]; then
    #     docker rmi $REGISTRY/$BACKEND_IMAGE
    # fi

    # Delete all Kubernetes resources
    kubectl delete -f ingress.yaml
    kubectl delete -f server_manifests/client_deployment.yaml
    kubectl delete -f client_service.yaml
    # kubectl delete -f server_manifests/backend_deployment.yaml
    # kubectl delete -f backend_service.yaml
    kubectl delete -f pvc.yaml
}

run() {
    # Cleanup existing resources
    cleanup

    # Start the Docker registry
    docker run -d -p 5000:5000 --restart=always --name registry registry:2

    # Build and push the client Docker image
    if [[ "$(docker images -q $CLIENT_IMAGE 2> /dev/null)" == "" ]]; then
        echo "Building the client Docker image..."
        cd ../client
        docker build -t $CLIENT_IMAGE .
        cd ../dummy
    fi

    # # Build and push the backend Docker image
    # if [[ "$(docker images -q $BACKEND_IMAGE 2> /dev/null)" == "" ]]; then
    #     echo "Building the backend Docker image..."
    #     cd ../backend
    #     docker build -t $BACKEND_IMAGE .
    #     cd ../dummy
    # fi

    docker tag $CLIENT_IMAGE $REGISTRY/$CLIENT_IMAGE
    # docker tag $BACKEND_IMAGE $REGISTRY/$BACKEND_IMAGE

    docker push $REGISTRY/$CLIENT_IMAGE
    # docker push $REGISTRY/$BACKEND_IMAGE

    # Apply Kubernetes manifests
    kubectl apply -f ingress.yaml
    kubectl apply -f server_manifests/client_deployment.yaml
    kubectl apply -f client_service.yaml
    # kubectl apply -f server_manifests/backend_deployment.yaml
    # kubectl apply -f backend_service.yaml
    kubectl apply -f pvc.yaml

    # ./curl_pod/pipeline.sh
}

run
