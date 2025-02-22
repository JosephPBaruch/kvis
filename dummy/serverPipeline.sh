#!/bin/bash

# filepath: /Users/joseph.baruch/REPO/kvis/dummy/run_k3d.sh

# Variables
HOSTNAME="localhost:5000"
IP="127.0.0.1"
HOSTS_FILE="/etc/hosts"
CLIENT_IMAGE="kvis-client:latest"
BACKEND_IMAGE="kvis-backend:latest"
REGISTRY="localhost:5000"

run() {
    # LOADBALANCING FOR THE CORRECT PORTS


    if [[ "$(docker images -q $CLIENT_IMAGE 2> /dev/null)" == "" ]]; then
        echo "Building the client Docker image..."
        cd ../client
        docker build -t $CLIENT_IMAGE .
        cd ../dummy
    fi

    if [[ "$(docker images -q $BACKEND_IMAGE 2> /dev/null)" == "" ]]; then
        echo "Building the backend Docker image..."
        cd ../backend
        docker build -t $BACKEND_IMAGE .
        cd ../dummy
    fi

    docker tag $CLIENT_IMAGE $REGISTRY/$CLIENT_IMAGE
    docker tag $BACKEND_IMAGE $REGISTRY/$BACKEND_IMAGE

    docker push $REGISTRY/$CLIENT_IMAGE
    docker push $REGISTRY/$BACKEND_IMAGE

    kubectl apply -f ingress.yaml
    kubectl apply -f server_manifests/client_deployment.yaml
    kubectl apply -f client_service.yaml
    kubectl apply -f server_manifests/backend_deployment.yaml
    kubectl apply -f backend_service.yaml
    kubectl apply -f pvc.yaml

    # ./curl_pod/pipeline.sh
}

run
