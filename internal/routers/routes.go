package routers

import (
	"net/http"
)

func InitRoutes(server *http.ServeMux) {
	//Routers
	indexRouter := NewIndexRouter(server)

	//Initialization all routes
	indexRouter.InitRoutes()
}
