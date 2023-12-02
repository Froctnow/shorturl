package routers

import (
	"net/http"
	"shorturl/internal/handlers"
	"shorturl/internal/services"
)

type IndexRouter struct {
	server *http.ServeMux
}

func NewIndexRouter(server *http.ServeMux) *IndexRouter {
	indexRouter := &IndexRouter{server: server}

	return indexRouter
}

func (ir *IndexRouter) InitRoutes() {
	urlService := services.NewURLService()
	urlHandler := handlers.NewURLHandler(urlService)

	ir.server.HandleFunc("/", urlHandler.HandleRequest())
}
