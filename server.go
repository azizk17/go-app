package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber"
	"github.com/sujit-baniya/fiber-boilerplate/rest/middlewares"
	"github.com/techschool/myApp/config"
	db "github.com/techschool/myApp/db/sqlc"
	"github.com/techschool/myApp/token"
)

// Server serves HTTP requests for our banking service.
type Server struct {
	config     config.Config
	store      db.Store
	tokenMaker token.Maker
	router     *fiber.App
}

// NewServer creates a new HTTP server and set up routing.
func New(config config.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.DB.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	// if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
	// 	v.RegisterValidation("currency", validCurrency)
	// }

	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {

	api := server.router.Group("api").Use(middlewares.AuthApi())
	web := server.router.Group("")
	ApiRoutes(api)

	router := gin.Default()

	router.POST("/users", server.createUser)
	router.POST("/users/login", server.loginUser)

	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))
	authRoutes.POST("/accounts", server.createAccount)
	authRoutes.GET("/accounts/:id", server.getAccount)
	authRoutes.GET("/accounts", server.listAccounts)

	authRoutes.POST("/transfers", server.createTransfer)

	server.router = router
}

// Start runs the HTTP server on a specific address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
