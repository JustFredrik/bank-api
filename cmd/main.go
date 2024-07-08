package main

import (
	// Load env vars before all other packages
	_ "github.com/justfredrik/bank-api/internal/envloader"

	// Internal
	"github.com/justfredrik/bank-api/internal/auth"
	"github.com/justfredrik/bank-api/internal/db"
	"github.com/justfredrik/bank-api/internal/handlers"

	//External
	"github.com/gin-gonic/gin"
)

func init() {

	// Initialize Local mock DB with camt053 data
	if err := db.InitializeLocalMockData(); err != nil {
		panic(err)
	}

	// Initialize mock API Keys for Admin and user
	auth.CreateMockKeys()
}

func main() {

	router := gin.Default() // Default router uses middlewares Logger and Recovery

	{ // Declare Routes  ==================================================================

		router.GET("/ping", handlers.GetPing)

		// Only Admin can list all accounts
		router.GET("/accounts", auth.Authenticator(auth.ROLE_ADMIN), handlers.GetAccounts)

		// Endpoints that Require Account AUTH or admin AUTH
		accountAuthGroup := router.Group("/accounts")
		accountAuthGroup.Use(auth.Authenticator(auth.ROLE_ACCOUNT))
		{ // Routes
			accountAuthGroup.GET("/:accountId", handlers.GetAccount)
			accountAuthGroup.GET("/:accountId/transactions", handlers.GetAccount)
			accountAuthGroup.GET("/:accountId/transactions/:transactionId", handlers.GetAccount)
		}
	}

	router.Run()
}
