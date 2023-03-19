package api

import (
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine, srv *Server) {
	routes := router.Group("/api")
	{

		// health check
		routes.GET("/health", healthy)

		// auth routes
		authRoute := routes.Group("/auth")
		{
			authRoute.GET("/current_user", srv.getCurrentUserBySession)
			authRoute.POST("/login", srv.loginUser)
			authRoute.POST("/register", srv.CreateUser)
			// authRoute.POST("/logout", srv.Logout)
		}

		// user routes
		userRoute := routes.Group("/user")
		{
			// userRoute.POST("/", srv.CreateUser)
			userRoute.GET("/", srv.GetUsers)
			userRoute.GET("/:id", srv.GetUserByID)
			// userRoute.GET("/:email", srv.GetUserByEmail)

		}
	}
}
