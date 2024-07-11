package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/justfredrik/bank-api/internal/db"
)

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

func GetPing(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func GetAccount(c *gin.Context) {

	// validate accountId param format
	id, err := validateAccountIdParam(c)
	if err != nil {
		return
	}

	// Fetch Account from mock db
	acc, err := db.DB.GetAccount(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Not Found", "message": "account not found"})
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
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "Not Found",
			"message": "unable to find accounts",
		})
		return
	}

	c.JSON(http.StatusOK, accounts)
}

func GetTransaction(c *gin.Context) {

	// validate accountId param format
	accountId, err := validateAccountIdParam(c)
	if err != nil {
		// This should technically be unreachable since this has already been validated in the AUTH step.
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error", "message": "the server was uanble to validate the accountId"})
		return
	}

	// Fetch Transaction from mock db
	transaction, err := db.DB.GetAccountTransaction(accountId, c.Param("transactionRef"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Not Found", "message": "transaction not found"})
		return
	}

	// Return valid transaction data
	c.JSON(http.StatusOK, *transaction)

}

func GetTransactions(c *gin.Context) {
	// No support for pagination but would be good to have if in real prod
	// validate accountId param format
	accountId, err := validateAccountIdParam(c)
	if err != nil {
		// This should technically be unreachable since this has already been validated in the AUTH step.
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error", "message": "the server was uanble to validate the accountId"})
		return
	}

	// Fetch transactions from mock db
	transactions, err := db.DB.GetAccountTransactions(accountId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error", "message": "the server was uanble to fetch the transactions"})
		return
	}

	// Return valid transactions data
	c.JSON(http.StatusOK, transactions)

}
