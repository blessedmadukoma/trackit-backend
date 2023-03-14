package api

import (
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine, srv *Server) {
	routes := router.Group("/api")
	{

		// user routes
		userRoute := routes.Group("/user")
		{
			userRoute.POST("/", srv.CreateUser)
			userRoute.GET("/", srv.GetUsers)
			userRoute.GET("/:id", srv.GetUserByID)
			// userRoute.GET("/:email", srv.GetUserByEmail)

		}
	}
}
