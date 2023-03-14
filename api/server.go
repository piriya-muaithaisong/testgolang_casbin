package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/piriya-muaithaisong/testgolang_casbin/db/sqlc"
	"github.com/piriya-muaithaisong/testgolang_casbin/utils"
)

type Server struct {
	config utils.Config
	store  db.Store
	router *gin.Engine
	// s
}

func NewServer(config utils.Config, store db.Store) (*Server, error) {
	// tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)

	// if err != nil {
	// 	return nil, fmt.Errorf("cannot create token maker: %w", err)
	// }
	server := &Server{
		store: store,
		// tokenMaker: tokenMaker,
		config: config,
	}

	// if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
	// 	v.RegisterValidation("currency", validCurrency)
	// }

	server.setupRouter()

	return server, nil
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func (server *Server) setupRouter() {
	router := gin.Default()

	// router.POST("/users", server.createUser)
	// router.POST("users/login", server.loginUser)

	// router.POST("/token/renew_access", server.renewAccessToken)

	// //require auth
	// authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))

	// authRoutes.POST("/accounts", server.createAccount)
	// authRoutes.GET("/accounts/:id", server.getAccount)
	// authRoutes.GET("/accounts", server.listAccount)

	// authRoutes.POST("/transfers", server.createTransfer)

	server.router = router
}
