package routes

import (
	"backend/handlers"
	"net/http"
)

func RegisterRoutes() {
	http.HandleFunc("/", handlers.HelloHandler)
	http.HandleFunc("/pods", handlers.PodsHandler)
}
