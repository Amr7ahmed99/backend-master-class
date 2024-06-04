package apis

import (
	"backend-master-class/apis/middlewares"
	db "backend-master-class/db/sqlc"
	"backend-master-class/token"
	"backend-master-class/util"
	"backend-master-class/validators"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// Server serves HTTP requests for our banking service.
type Server struct {
	Store      db.Store
	Router     *gin.Engine
	TokenMaker token.Maker
	Config     util.Config
}

// NewServer creates a new HTTP server and setup routing.
func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewJWTMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		Store:      store,
		TokenMaker: tokenMaker,
		Config:     config,
	}

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validators.ValidCurrency)
	}

	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()
	authRoutes := router.Group("/")
	{
		authRoutes.Use(middlewares.AuthMiddleware(server.TokenMaker))

		authRoutes.POST("/accounts", server.createAccount)
		authRoutes.GET("/accounts/:id", server.getAccount)
		authRoutes.GET("/accounts", server.listAccounts)
		authRoutes.PUT("/accounts/:id", server.updateAccount)

		authRoutes.POST("/transfers", server.createTransfer)

		authRoutes.GET("/users", server.listUsers)
		authRoutes.GET("/users/:username", server.getUser)
	}

	router.POST("/users", server.createUser)
	router.POST("/users/login", server.loginUser)

	server.Router = router
}

// Start runs the HTTP server on a specific address
func (server *Server) Start(address string) error {
	return server.Router.Run(address)
}

func (server *Server) errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
