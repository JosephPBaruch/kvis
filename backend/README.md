# kvis-backend

## API 

### kubectl get endpoints

    /pods
    /nodes
    /deployments
    /services
    /namespaces
    /configmaps
    /pvc
    /ingress
    /endpoints

### kubectl info endpoints

These will display information from kubectl describe, logs, events

    /pods/{name_id}
        Complete
    /nodes/{name_id}
    /deployments/{name_id}
    /services/{name_id}
        Complete
    /namespaces/{name_id}
    /configmaps/{name_id}
    /pvc/{name_id}
    /ingress/{name_id}
        Complete
    /endpoints/{name_id}
        / Complete 