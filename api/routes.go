package api

import (
	"net/http"

	"github.com/TonyGLL/erp_backend/docs"
	"github.com/TonyGLL/erp_backend/middlewares"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (server *Server) SetupRoutes(version string) http.Handler {
	r := gin.Default()

	docs.SwaggerInfo.Title = "Swagger ERP Documentation"
	docs.SwaggerInfo.Description = "This is an ERP Backend with Golang."
	docs.SwaggerInfo.Version = version
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http"}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	/* CORS */
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://foo.com"},
		AllowMethods:     []string{"PUT", "PATCH", "OPTIONS", "GET", "POST"},
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

		/* ROLES */
		v1.GET("/roles", server.getRoles)

		/* MODULES */

		/* PAYROLLS */

		/* EMAIL */
		v1.POST("/email", server.sendEmail)
	}

	return r
}
