package main

import (
	"net/http"
	"shorturl/internal/routers"
)

func main() {
	server := http.NewServeMux()

	routers.InitRoutes(server)

	_ = http.ListenAndServe(":8080", server)
}
