package main

import (
	"backend/routes"
	"fmt"
	"log"
	"net/http"
)

func main() {
	routes.RegisterRoutes()
	fmt.Println("Server is running on port 8082...")
	log.Fatal(http.ListenAndServe(":8082", nil))
}
