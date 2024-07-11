// package handlers provides handler functions linking the endpoints in the router to other internal systems.
package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/justfredrik/bank-api/internal/db"
)

// validateAccountIdParam makes sure that the accountId is a uint64.
func validateAccountIdParam(c *gin.Context) (uint64, error) {
	id, err := strconv.ParseUint(c.Param("accountId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": "id is not a positive integer",
		})
		c.Abort()
	}
	return id, err
}

// GetPing is a gin Handler that returns a pong resonse to the requester.
func GetPing(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

// GetAccount is a gin Handler that retrieves an account based on the accountId parameter.
func GetAccount(c *gin.Context) {

	id, err := validateAccountIdParam(c)
	if err != nil {
		return
	}

	acc, err := db.DB.GetAccount(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Not Found", "message": "account not found"})
		return
	}

	c.JSON(http.StatusOK, *acc)

}

// GetAccounts is a gin Handler that returns a list of accounts to the requester.
func GetAccounts(c *gin.Context) {
	// No support for pagination but would be good to have if in real prod

	accounts, err := db.DB.GetAccounts(0, 0)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "Not Found",
			"message": "unable to find accounts",
		})
		return
	}

	c.JSON(http.StatusOK, accounts)
}

// GetTransaction is a gin Handler that returns a specific account transaction to the requester.
func GetTransaction(c *gin.Context) {

	accountId, err := validateAccountIdParam(c)
	if err != nil {
		// This should technically be unreachable since this has already been validated in the AUTH step.
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error", "message": "the server was uanble to validate the accountId"})
		return
	}

	transaction, err := db.DB.GetAccountTransaction(accountId, c.Param("transactionRef"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Not Found", "message": "transaction not found"})
		return
	}

	c.JSON(http.StatusOK, *transaction)

}

// GetTransactions returns a list of transactions associated with an account.
func GetTransactions(c *gin.Context) {
	// No support for pagination but would be good to have if in real prod

	accountId, err := validateAccountIdParam(c)
	if err != nil {
		// This should technically be unreachable since this has already been validated in the AUTH step.
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error", "message": "the server was uanble to validate the accountId"})
		return
	}

	transactions, err := db.DB.GetAccountTransactions(accountId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error", "message": "the server was uanble to fetch the transactions"})
		return
	}

	c.JSON(http.StatusOK, transactions)

}
