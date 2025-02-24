
#!/bin/bash

echo "Deleting specific deployments..."

kubectl delete -f /Users/joseph.baruch/REPO/personal/kvis/dummy/server_manifests/backend_deployment.yaml
kubectl delete -f /Users/joseph.baruch/REPO/personal/kvis/dummy/server_manifests/client_deployment.yaml
kubectl delete -f /Users/joseph.baruch/REPO/personal/kvis/dummy/ingress.yaml
kubectl delete -f /Users/joseph.baruch/REPO/personal/kvis/dummy/client_service.yaml
kubectl delete -f /Users/joseph.baruch/REPO/personal/kvis/dummy/backend_service.yaml
kubectl delete -f /Users/joseph.baruch/REPO/personal/kvis/dummy/pvc.yaml

echo "Specific deployments have been deleted."
