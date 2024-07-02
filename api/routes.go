package api

import (
	"net/http"

	"github.com/TonyGLL/erp_backend/middlewares"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func (server *Server) SetupRoutes() http.Handler {
	r := gin.Default()

	/* CORS */
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://foo.com"},
		AllowMethods:     []string{"PUT", "PATCH"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
	}))

	v1 := r.Group("/api/v1")
	{
		/* AUTH */
		v1.POST("/auth/login", server.AuthLogin)

		/* AUTH MIDDLEWARE */
		v1.Use(middlewares.AuthenticateMiddleware)

		/* USERS */
		v1.GET("/users", server.getUsers)
		v1.POST("/users", server.createUser)
		v1.GET("/users/:id", server.getUser)
		v1.PUT("/users/:id", server.updateUser)
		v1.DELETE("/users/:id", server.deleteUser)
		v1.GET("/users/download/csv", server.downloadUsersCSV)
	}

	return r
}
