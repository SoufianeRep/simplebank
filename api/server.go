package api

import (
	db "github.com/SoufianeRep/simplebank/db/sqlc"
	"github.com/gin-gonic/gin"
)

// Server serves HTTP requests for our bqnking service
type Server struct {
	store  *db.Store
	router *gin.Engine
}

// Creates a new http server
func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.GET("/accounts", server.listAccounts)
	router.GET("/accounts/:id", server.getAccount)
	router.POST("/accounts", server.createAccount)
	router.PUT("/accounts/:id", server.updateAccount)
	// router.DELETE("/accounts/:id", server.deleteAccount)

	server.router = router
	return server
}

// Start http server on the provided address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
