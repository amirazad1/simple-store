package api

import (
	"github.com/amirazad1/simple-store/service"
	"github.com/gin-gonic/gin"
)

type Server struct {
	store  service.Store
	router *gin.Engine
}

func NewServer(store service.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/products", server.createProduct)
	router.GET("/products/:id", server.getProduct)
	router.GET("/products", server.listProduct)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{
		"error": err.Error(),
	}
}
