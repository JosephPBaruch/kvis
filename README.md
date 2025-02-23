# kvis

## Overview

kvis is a tool designed to provide real-time information about the Kubernetes cluster running on your system. It consists of a backend written in Go and a frontend built with React and Vite.

## Backend

The backend is written in Go and is responsible for executing `kubectl` commands to gather information about the current state of the Kubernetes cluster. It exposes this information through a RESTful API.

### Features

- Executes `kubectl get pods` to retrieve information about the pods running in the cluster.
- Provides a RESTful API to serve this information to the frontend.

## Frontend

The frontend is built with React and Vite. It makes requests to the backend to fetch the current state of the Kubernetes cluster and displays this information in a user-friendly manner.

### Features

- Fetches pod information from the backend.
- Displays the pod information in a structured and helpful way.

## Getting Started

### Prerequisites

- Docker
- kubectl
- k3d (for running a local Kubernetes cluster)

### Running the Project

1. **Clone the repository:**

    ```sh
    git clone https://github.com/yourusername/kvis.git
    cd kvis
    ```

2. **Build and run the backend:**

    ```sh
    cd backend
    docker build -t kvis-backend .
    docker run -d -p 8080:8080 --name kvis-backend kvis-backend
    ```

3. **Build and run the frontend:**

    ```sh
    cd ../frontend
    docker build -t kvis-frontend .
    docker run -d -p 80:80 --name kvis-frontend kvis-frontend
    ```

4. **Run a local Kubernetes cluster using k3d:**

    ```sh
    k3d cluster create mycluster
    ```

5. **Access the frontend:**

    Open your browser and navigate to `http://localhost`.

## License

This project is licensed under the MIT License - see the [LICENSE](http://_vscodecontentref_/1) file for details.
