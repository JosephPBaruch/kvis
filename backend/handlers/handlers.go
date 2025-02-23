package handlers

import (
	"backend/control"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func DefaultHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("kvis-backend"))
}

func PodsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	pods, err := control.Get("pods")
	if err != nil {
		http.Error(w, "Failed to get pods", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(pods)
}

func PodByNameHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]

	details, err := control.GetPodDetails(id)
	if err != nil {
		http.Error(w, "Failed to get pod details", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(details)
}

func NodesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	nodes, err := control.Get("nodes")
	if err != nil {
		http.Error(w, "Failed to get nodes", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(nodes)
}

func NodeByNameHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]

	details, err := control.GetNodeDetails(id)
	if err != nil {
		http.Error(w, "Failed to get node details", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(details)
}

func DeploymentsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	print("DeploymentsHandler")

	deployments, err := control.Get("deployments")
	if err != nil {
		http.Error(w, error.Error(err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(deployments)
}

func DeploymentByNameHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]

	details, err := control.GetDeploymentDetails(id)
	if err != nil {
		http.Error(w, "Failed to get deployment details", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(details)
}

func ServicesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	services, err := control.Get("services")
	if err != nil {
		http.Error(w, "Failed to get services", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(services)
}

func ServiceByNameHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]

	details, err := control.GetServiceDetails(id)
	if err != nil {
		http.Error(w, "Failed to get service details", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(details)
}

func NamespacesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	namespaces, err := control.Get("namespaces")
	if err != nil {
		http.Error(w, "Failed to get namespaces", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(namespaces)
}

func NamespaceByNameHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]

	details, err := control.GetNamespaceDetails(id)
	if err != nil {
		http.Error(w, "Failed to get namespace details", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(details)
}

func ConfigMapsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	configmaps, err := control.Get("configmaps")
	if err != nil {
		http.Error(w, "Failed to get configmaps", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(configmaps)
}

func ConfigMapByNameHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]

	details, err := control.GetConfigMapDetails(id)
	if err != nil {
		http.Error(w, "Failed to get configmap details", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(details)
}

func PVCHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	pvcs, err := control.Get("pvc")
	if err != nil {
		http.Error(w, "Failed to get PVCs", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(pvcs)
}

func PVCByNameHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]

	details, err := control.GetPVCDetails(id)
	if err != nil {
		http.Error(w, "Failed to get PVC details", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(details)
}

func IngressHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	ingresses, err := control.Get("ingress")
	if err != nil {
		http.Error(w, error.Error(err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ingresses)
}

func IngressByNameHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]

	details, err := control.GetIngressDetails(id)
	if err != nil {
		http.Error(w, "Failed to get ingress details", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(details)
}

func EndpointsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	endpoints, err := control.Get("endpoints")
	if err != nil {
		http.Error(w, "Failed to get endpoints", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(endpoints)
}

func EndpointByNameHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]

	details, err := control.GetEndpointDetails(id)
	if err != nil {
		http.Error(w, "Failed to get endpoint details", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(details)
}
