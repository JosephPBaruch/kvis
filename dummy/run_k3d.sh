#!/bin/bash

# Check if k3d is installed
if ! command -v k3d &> /dev/null
then
    echo "k3d could not be found, installing k3d..."
    curl -s https://raw.githubusercontent.com/rancher/k3d/main/install.sh | bash
else
    echo "k3d is already installed."
fi

# Check if the cluster is already running
if k3d cluster list | grep -q 'mycluster'
then
    echo "k3d cluster 'mycluster' is already running. Deleting the existing cluster..."
    k3d cluster delete mycluster
fi

# Create a new k3d cluster
echo "Creating k3d cluster..."
k3d cluster create mycluster

# Check if the cluster is running
if k3d cluster list | grep -q 'mycluster'
then
    echo "k3d cluster 'mycluster' is up and running."
else
    echo "Failed to create k3d cluster."
    exit 1
fi

echo "You can now use kubectl to interact with your k3d cluster."

kubectl apply -f /Users/joseph.baruch/REPO/kvis/dummy/deployment.yaml