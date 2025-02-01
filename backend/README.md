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
        Complete (kind of) This needs to be revised in the future
    /deployments/{name_id}
        Complete
    /services/{name_id}
        Complete
    /namespaces/{name_id}
        Complete
    /configmaps/{name_id}
        Complete
    /pvc/{name_id}
        Complete
    /ingress/{name_id}
        Complete
    /endpoints/{name_id}
        Complete

# Backend Server

## Running the Server

To run the server with HTTPS, you need to generate SSL certificates. You can use OpenSSL to generate a self-signed certificate for testing purposes.

### Generate SSL Certificates

```sh
openssl req -x509 -newkey rsa:4096 -keyout server.key -out server.crt -days 365 -nodes
```

### Run the Server

```sh
go run main.go
```

The server will be available at https://localhost:8082.