#!/bin/bash

IMAGE="curl-image"
REGISTRY="k3d-local-registry:5001"

docker build -t $IMAGE .

docker tag $IMAGE $REGISTRY/$IMAGE

docker push $REGISTRY/$IMAGE

kubectl apply -f curl_pod.yaml


# Exec inside the pod
# kubectl exec -it curl-image-pod -- /bin/sh

# Make a cURL request to the backend via the backend service
# curl http://kvis-backend.default.svc.cluster.local:8082