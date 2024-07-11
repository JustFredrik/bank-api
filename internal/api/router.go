package api

import (
	"github.com/gin-gonic/gin"
	"github.com/justfredrik/bank-api/internal/api/handlers"
	"github.com/justfredrik/bank-api/internal/auth"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default() // Default router uses middlewares Logger and Recovery

	// ====================================================================================
	{ // Declare Routes

		router.GET("/ping", auth.Authenticator(auth.ROLE_ANY), handlers.GetPing)

		// Only Admin can list all accounts
		router.GET("/accounts", auth.Authenticator(auth.ROLE_ADMIN), handlers.GetAccounts)

		// Endpoints that Require Account AUTH or admin AUTH
		accountAuthGroup := router.Group("/accounts")
		accountAuthGroup.Use(auth.Authenticator(auth.ROLE_ACCOUNT))
		{ // Routes
			accountAuthGroup.GET("/:accountId", handlers.GetAccount)
			accountAuthGroup.GET("/:accountId/transactions", handlers.GetTransactions)
			accountAuthGroup.GET("/:accountId/transactions/:transactionRef", handlers.GetTransaction)
		}
	}
	return router
}
