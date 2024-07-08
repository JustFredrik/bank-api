package main

import (
	// Load env vars before all other packages
	_ "github.com/justfredrik/bank-api/internal/envloader"

	// Internal
	"github.com/justfredrik/bank-api/internal/db"
	"github.com/justfredrik/bank-api/internal/handlers"

	//External
	"github.com/gin-gonic/gin"
)

func main() {

	if err := db.InitializeLocalMockData(); err != nil {
		panic(err)
	}

	router := gin.Default()

	router.GET("/accounts", handlers.GetAccounts)
	router.GET("/accounts/:id", handlers.GetAccount)

	// jsonBytes, _ := json.MarshalIndent(camtDoc.BankStatement, "", "	")
	// fmt.Print(string(jsonBytes))
	// fmt.Print(initializer.Bob)

	router.Run()
}
