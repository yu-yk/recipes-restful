package api

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

// Server contains a db and router
type Server struct {
	db     *sql.DB
	router *gin.Engine
}

// NewServer returns an api server instance
func NewServer(db *sql.DB) *Server {
	return &Server{
		db:     db,
		router: gin.Default(),
	}
}
