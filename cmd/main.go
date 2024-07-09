package main

import (
	// Load env vars before all other packages
	_ "github.com/justfredrik/bank-api/internal/envloader"

	// Internal packages
	"github.com/justfredrik/bank-api/internal/api"
	"github.com/justfredrik/bank-api/internal/auth"
	"github.com/justfredrik/bank-api/internal/db"
)

func init() {

	// ========================================================
	// Initialize Local mock DB with camt053 data
	// ========================================================
	if err := db.InitializeLocalMockData(); err != nil {
		panic(err)
	}

	// ========================================================
	// Initialize mock API Keys for Admin and user for testing
	// ========================================================
	auth.CreateMockKeys()
}

func main() {
	router := api.SetUpRouter()
	router.Run()
}
