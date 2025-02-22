package routes

import (
	"backend/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func Routes() {
	r := mux.NewRouter()

	// Add CORS middleware
	r.Use(mux.CORSMethodMiddleware(r))
	r.Use(corsMiddleware)

	r.HandleFunc("backend/", handlers.DefaultHandler)
	r.HandleFunc("backend/pods", handlers.PodsHandler)
	r.HandleFunc("backend/pods/{id}", handlers.PodByNameHandler)
	r.HandleFunc("backend/nodes", handlers.NodesHandler)
	r.HandleFunc("backend/nodes/{id}", handlers.NodeByNameHandler)
	r.HandleFunc("backend/deployments", handlers.DeploymentsHandler)
	r.HandleFunc("backend/deployments/{id}", handlers.DeploymentByNameHandler)
	r.HandleFunc("backend/services", handlers.ServicesHandler)
	r.HandleFunc("backend/services/{id}", handlers.ServiceByNameHandler)
	r.HandleFunc("backend/namespaces", handlers.NamespacesHandler)
	r.HandleFunc("backend/namespaces/{id}", handlers.NamespaceByNameHandler)
	r.HandleFunc("backend/configmaps", handlers.ConfigMapsHandler)
	r.HandleFunc("backend/configmaps/{id}", handlers.ConfigMapByNameHandler)
	r.HandleFunc("backend/pvc", handlers.PVCHandler)
	r.HandleFunc("backend/pvc/{id}", handlers.PVCByNameHandler)
	r.HandleFunc("backend/ingress", handlers.IngressHandler)
	r.HandleFunc("backend/ingress/{id}", handlers.IngressByNameHandler)
	r.HandleFunc("backend/endpoints", handlers.EndpointsHandler)
	r.HandleFunc("backend/endpoints/{id}", handlers.EndpointByNameHandler)

	http.Handle("/", r)
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}
