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

# Check if the local registry is already running
if k3d registry list | grep -q 'k3d-local-registry'
then
    echo "k3d local registry 'k3d-local-registry' is already running. Deleting the existing registry..."
    k3d registry delete k3d-local-registry
fi

k3d registry create local-registry --port 5001

#!/bin/bash

# Define the hostname and IP mapping
HOSTNAME="k3d-local-registry"
IP="127.0.0.1"
HOSTS_FILE="/etc/hosts"

# Check if the hostname already exists in /etc/hosts
if grep -q "$HOSTNAME" "$HOSTS_FILE"; then
  echo "The hostname '$HOSTNAME' already exists in $HOSTS_FILE."
else
  # Add the hostname to /etc/hosts
  echo "Adding '$HOSTNAME' to $HOSTS_FILE..."
  echo "$IP $HOSTNAME" | sudo tee -a "$HOSTS_FILE" > /dev/null

  # Confirm the addition
  if grep -q "$HOSTNAME" "$HOSTS_FILE"; then
    echo "Successfully added '$HOSTNAME' to $HOSTS_FILE."
  else
    echo "Failed to add '$HOSTNAME' to $HOSTS_FILE. Please check permissions."
  fi
fi

k3d cluster create mycluster --registry-use k3d-local-registry:5001

# Check if the cluster is running
if k3d cluster list | grep -q 'mycluster'
then
    echo "k3d cluster 'mycluster' is up and running."
else
    echo "Failed to create k3d cluster."
    exit 1
fi

echo "You can now use kubectl to interact with your k3d cluster."

docker tag kvis-client:latest k3d-local-registry:5001/kvis-client:latest

docker push k3d-local-registry:5001/kvis-client:latest
# Check with
# curl http://localhost:5001/v2/_catalog

kubectl apply -f client_deployment.yaml
kubectl apply -f client_service.yaml
kubectl apply -f deployment.yaml