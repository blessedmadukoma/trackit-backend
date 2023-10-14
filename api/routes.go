package api

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Routes(router *gin.Engine, srv *Server) {
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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

			userRoute.POST("/expense", srv.CreateExpense)
			userRoute.GET("/expense", srv.GetExpenses)
		}

		// expense routes
		expenseRoute := routes.Group("/expense")
		{
			expenseRoute.POST("/", srv.CreateExpense)
			expenseRoute.GET("/", srv.GetExpenses)
			// expenseRoute.GET("/:id", srv.GetExpenseByID)
			// expenseRoute.PUT("/:id", srv.UpdateExpense)
			// expenseRoute.DELETE("/:id", srv.DeleteExpense)
		}
	}
}
