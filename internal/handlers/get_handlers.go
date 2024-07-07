package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/justfredrik/bank-api/internal/db"
)

func GetPing(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func GetAccount(c *gin.Context) {

	// validate id param format
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "id is not a positive integer",
		})
	}

	// Fetch Account from mock db
	acc, err := db.DB.GetAccount(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{})
		return
	}

	// Return valid account data
	c.JSON(http.StatusOK, *acc)

}

func GetAccounts(c *gin.Context) {

	// No support for pagination but would be good to have if in real prod

	// Fetch Accounts from mock db
	accounts, err := db.DB.GetAccounts(0, 0)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{})
		return
	}

	c.JSON(http.StatusOK, accounts)
}
