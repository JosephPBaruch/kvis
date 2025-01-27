package routes

import (
	"backend/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func Routes() {
	r := mux.NewRouter()
	r.HandleFunc("/", handlers.DefaultHandler)
	r.HandleFunc("/pods", handlers.PodsHandler)
	r.HandleFunc("/pods/{id}", handlers.PodByNameHandler)
	r.HandleFunc("/nodes", handlers.NodesHandler)
	r.HandleFunc("/nodes/{id}", handlers.NodeByNameHandler)
	r.HandleFunc("/deployments", handlers.DeploymentsHandler)
	r.HandleFunc("/deployments/{id}", handlers.DeploymentByNameHandler)
	r.HandleFunc("/services", handlers.ServicesHandler)
	r.HandleFunc("/services/{id}", handlers.ServiceByNameHandler)
	r.HandleFunc("/namespaces", handlers.NamespacesHandler)
	r.HandleFunc("/namespaces/{id}", handlers.NamespaceByNameHandler)
	r.HandleFunc("/configmaps", handlers.ConfigMapsHandler)
	r.HandleFunc("/configmaps/{id}", handlers.ConfigMapByNameHandler)
	r.HandleFunc("/pvc", handlers.PVCHandler)
	r.HandleFunc("/pvc/{id}", handlers.PVCByNameHandler)
	r.HandleFunc("/ingress", handlers.IngressHandler)
	r.HandleFunc("/ingress/{id}", handlers.IngressByNameHandler)
	r.HandleFunc("/endpoints", handlers.EndpointsHandler)
	r.HandleFunc("/endpoints/{id}", handlers.EndpointByNameHandler)

	http.Handle("/", r)
}
