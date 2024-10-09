package api

import (
	"fmt"
	"github.com/amirazad1/simple-store/service"
	"github.com/amirazad1/simple-store/token"
	"github.com/amirazad1/simple-store/util"
	"github.com/gin-gonic/gin"
)

type Server struct {
	config     util.Config
	store      service.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

func NewServer(config util.Config, store service.Store) (*Server, error) {
	tokenMaker, err := token.NewJWTMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}
	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()

	router.POST("/users", server.createUser)
	router.POST("/users/login", server.loginUser)

	auth := router.Group("/").Use(authMiddleware(server.tokenMaker))
	auth.POST("/products", server.createProduct)
	auth.GET("/products/:id", server.getProduct)
	auth.GET("/products", server.listProduct)

	auth.POST("/sales", server.createSale)

	auth.GET("/users/:name", server.getUser)

	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{
		"error": err.Error(),
	}
}
