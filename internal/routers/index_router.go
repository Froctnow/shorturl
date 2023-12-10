package routers

import (
	"net/http"
	"shorturl/internal/handlers"
	"shorturl/internal/services"
	"shorturl/internal/storage"
)

type IndexRouter struct {
	server *http.ServeMux
}

func NewIndexRouter(server *http.ServeMux) *IndexRouter {
	indexRouter := &IndexRouter{server: server}

	return indexRouter
}

func (ir *IndexRouter) InitRoutes() {
	ir.initURLRoutes()
}

func (ir *IndexRouter) initURLRoutes() {
	storageInstance := storage.NewStorage()
	urlRepository := storageInstance.URLRepository
	urlService := services.NewURLService(urlRepository)
	urlHandler := handlers.NewURLHandler(urlService)

	ir.server.HandleFunc("/", urlHandler.HandleRequest())
}
