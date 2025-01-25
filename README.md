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
- Go (for running the backend)

### Running the Project

1. **Clone the repository:**

    ```sh
    git clone https://github.com/yourusername/kvis.git
    cd kvis
    ```

2. **Create a symbolic link to the [kvis.sh](http://_vscodecontentref_/2) script:**

    ```sh
    ./link.sh
    ```

3. **Run the project:**

    ```sh
    kvis
    ```

4. **Run the project with the [dummy](http://_vscodecontentref_/3) script:**

    ```sh
    kvis dummy
    ```

### Running the Backend and Frontend Separately

#### Backend

To run the backend separately:

1. **Navigate to the backend directory:**

    ```sh
    cd backend
    ```

2. **Run the backend:**

    ```sh
    go run main.go
    ```

#### Frontend

To run the frontend separately:

1. **Navigate to the client directory:**

    ```sh
    cd client
    ```

2. **Run the frontend:**

    ```sh
    ./run_client.sh
    ```

### Running a Local Kubernetes Cluster Using k3d

1. **Run the [run_k3d.sh](http://_vscodecontentref_/4) script:**

    ```sh
    ./dummy/run_k3d.sh
    ```

2. **Apply the deployment:**

    ```sh
    kubectl apply -f /Users/joseph.baruch/REPO/kvis/dummy/deployment.yaml
    ```

### Accessing the Frontend

Open your browser and navigate to [http://localhost:8081](http://_vscodecontentref_/5).

## License

This project is licensed under the MIT License - see the [LICENSE](http://_vscodecontentref_/6) file for details.