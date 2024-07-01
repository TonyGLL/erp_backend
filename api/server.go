package api

import (
	"net/http"
	"time"

	db "github.com/TonyGLL/erp_backend/db/sqlc"
	"github.com/gin-gonic/gin"
)

type Server struct {
	port  string
	store db.Store
}

func NewServer(store db.Store, port string) *http.Server {
	NewServer := &Server{
		store: store,
		port:  port,
	}

	// Declare Server config
	server := &http.Server{
		Addr:         port,
		Handler:      NewServer.SetupRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
