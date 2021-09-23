package main

import (
	"fmt"
	"log"
	"net/http"
	"web/server/routes"
)

func main() {
	port := fmt.Sprintf(":%d", 8006)
	routes.LoadRoutes()
	log.Fatal(http.ListenAndServe(port, nil))
}
