package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/justfredrik/bank-api/internal/auth"
	"github.com/justfredrik/bank-api/internal/db"
	"github.com/stretchr/testify/assert"
)

func setUpTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	return SetUpRouter()
}

func jsonContains(jsonBytes []byte, requiredData map[string]string) (bool, error) {
	var parsedStrings map[string]string
	var parsedInter map[string]interface{}
	json.Unmarshal(jsonBytes, &parsedStrings)
	json.Unmarshal(jsonBytes, &parsedInter)

	for requiredKey, requiredValue := range requiredData {
		parsedValueS, firstOk := parsedStrings[requiredKey]
		_, secondOk := parsedInter[requiredKey]
		if firstOk && secondOk {
			// required value needs to be "" or match with response
			if (requiredValue == "") || (parsedValueS == requiredValue) {
				continue // move on to next requirement
			}
			return false, errors.New("value in key: " + requiredKey + ", does not match.\nExpected: " + requiredValue + "\nActual: " + parsedValueS + "\n")
		}
		return false, errors.New("does not contain key: " + requiredKey)
	}
	return true, nil // Matched all required keys and values without early return
}

type TestRequest struct {
	testName     string
	requestType  string
	endpoint     string
	headers      map[string]string
	expectedCode int
	expectedBody map[string]string
}

type errorResponse struct {
	message string `json: "message"`
	error   int    `json: "error"`
}

type PongResponse struct {
	Message string `json:"message"`
}

func setup() {
	if err := godotenv.Load("../../.env"); err != nil {
		panic(errors.New("Test Setup Failed. failed to load .env variables with error:" + err.Error()))
	}
	if err2 := db.InitializeLocalMockData(); err2 != nil {
		panic(errors.New("Test Setup Failed. failed to load mock DB with error:" + err2.Error()))
	}
	isSetup = true
}

var isSetup bool = false

func testReqests(t *testing.T, router *gin.Engine, tests []TestRequest) {

	if !isSetup {
		setup()
	}

	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {

			// Create Request
			req, _ := http.NewRequest(test.requestType, test.endpoint, nil)

			// Populate Headers
			for key, value := range test.headers {
				req.Header.Set(key, value)
			}

			// Serve the Request to the router
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			// Validate Response Status
			assert.Equal(t, test.expectedCode, w.Code, "Incorrect Response Status")

			// Validate Response Format
			bodyIsOK, err := jsonContains(w.Body.Bytes(), test.expectedBody)
			assert.Equal(t, true, bodyIsOK, err)

		})
	}
}

// TestAccounts tests GET requests to the /ping endpoint.
func TestPing(t *testing.T) {

	apiKey := auth.NewAPIKey(auth.ROLE_ACCOUNT, 1337)
	token := apiKey.Token()

	tests := []TestRequest{
		{
			testName:     "Valid API Key",
			requestType:  "GET",
			endpoint:     "/ping",
			headers:      map[string]string{"Authorization": "Bearer " + token},
			expectedCode: http.StatusOK,
			expectedBody: map[string]string{"message": "pong"},
		},
		{
			testName:     "Malformed header",
			requestType:  "GET",
			endpoint:     "/ping",
			headers:      map[string]string{"Authorization": "Beer " + token},
			expectedCode: http.StatusUnauthorized,
			expectedBody: map[string]string{"error": "Unauthorized", "message": "Malformed Authorization header"},
		},
		{
			testName:     "Missing header",
			requestType:  "GET",
			endpoint:     "/ping",
			headers:      map[string]string{"ThisIsWrong": "Fox " + token},
			expectedCode: http.StatusUnauthorized,
			expectedBody: map[string]string{"error": "Unauthorized", "message": "Missing Authorization header"},
		},
	}

	testReqests(t, setUpTestRouter(), tests)
}

// TestAccounts tests GET requests to the /accounts endpoint.
// This endpoint returns a list of accounts.
func TestAccounts(t *testing.T) {

	adminToken := auth.NewAPIKey(auth.ROLE_ADMIN, 1337).Token()
	accountToken := auth.NewAPIKey(auth.ROLE_ACCOUNT, 54400001111).Token()

	getTests := []TestRequest{
		{
			testName:     "Valid API Key",
			requestType:  "GET",
			endpoint:     "/accounts",
			headers:      map[string]string{"Authorization": "Bearer " + adminToken},
			expectedCode: http.StatusOK,
			expectedBody: map[string]string{"accounts": ""},
		},
		{
			testName:     "Unauthorized API Key",
			requestType:  "GET",
			endpoint:     "/accounts",
			headers:      map[string]string{"Authorization": "Bearer " + accountToken},
			expectedCode: http.StatusUnauthorized,
			expectedBody: map[string]string{"error": "Unauthorized", "message": "Your API key is not authorized to access the requested resource"},
		},
		{
			testName:     "Malformed header",
			requestType:  "GET",
			endpoint:     "/accounts",
			headers:      map[string]string{"Authorization": "Beer " + accountToken},
			expectedCode: http.StatusUnauthorized,
			expectedBody: map[string]string{"error": "Unauthorized", "message": "Malformed Authorization header"},
		},
		{
			testName:     "Missing header",
			requestType:  "GET",
			endpoint:     "/accounts",
			headers:      map[string]string{"ThisIsWrong": "Fox " + accountToken},
			expectedCode: http.StatusUnauthorized,
			expectedBody: map[string]string{"error": "Unauthorized", "message": "Missing Authorization header"},
		},
	}
	testReqests(t, setUpTestRouter(), getTests)

}

// TestAccountsAccountId tests GET requests to the /accounts/:accountId endpoint.
// This endpoint returns a specific account.
func TestAccountsAccountId(t *testing.T) {

	adminToken := (auth.NewAPIKey(auth.ROLE_ADMIN, 1337)).Token()
	accountToken := auth.NewAPIKey(auth.ROLE_ACCOUNT, 54400001111).Token()
	randomToken := auth.NewAPIKey(auth.ROLE_ACCOUNT, 1337).Token()

	getTests := []TestRequest{
		{
			testName:     "Valid Admin API Key",
			requestType:  "GET",
			endpoint:     "/accounts/54400001111",
			headers:      map[string]string{"Authorization": "Bearer " + adminToken},
			expectedCode: http.StatusOK,
			expectedBody: map[string]string{"account": "", "balances": ""},
		},
		{
			testName:     "Correct Account API Key",
			requestType:  "GET",
			endpoint:     "/accounts/54400001111",
			headers:      map[string]string{"Authorization": "Bearer " + accountToken},
			expectedCode: http.StatusOK,
			expectedBody: map[string]string{"account": "", "balances": ""},
		},
		{
			testName:     "Wrong Account API Key",
			requestType:  "GET",
			endpoint:     "/accounts/54400001111",
			headers:      map[string]string{"Authorization": "Bearer " + randomToken},
			expectedCode: http.StatusUnauthorized,
			expectedBody: map[string]string{"error": "Unauthorized", "message": "Your API key is not authorized to access the requested resource"},
		},
	}

	testReqests(t, setUpTestRouter(), getTests)
}

// TestAccountTransactions tests GET requests to the /accounts/:accountId/transactions endpoint.
// This endpoint returns a list of account transactions.
func TestAccountTransactions(t *testing.T) {

	adminToken := auth.NewAPIKey(auth.ROLE_ADMIN, 1337).Token()
	accountToken := auth.NewAPIKey(auth.ROLE_ACCOUNT, 54400001111).Token()
	randomToken := auth.NewAPIKey(auth.ROLE_ACCOUNT, 1337).Token()

	expectedOKBody := map[string]string{
		"transactions": "",
		"totalCount":   "",
		"page":         "",
		"perPage":      "",
	}

	expectedUnauthorizedBody := map[string]string{
		"error":   "Unauthorized",
		"message": "Your API key is not authorized to access the requested resource",
	}

	getTests := []TestRequest{
		{
			testName:     "Valid API Key (Admin)",
			requestType:  "GET",
			endpoint:     "/accounts/54400001111/transactions",
			headers:      map[string]string{"Authorization": "Bearer " + adminToken},
			expectedCode: http.StatusOK,
			expectedBody: expectedOKBody,
		},
		{
			testName:     "Valid API Key (Account Owner)",
			requestType:  "GET",
			endpoint:     "/accounts/54400001111/transactions",
			headers:      map[string]string{"Authorization": "Bearer " + accountToken},
			expectedCode: http.StatusOK,
			expectedBody: expectedOKBody,
		},
		{
			testName:     "Unauthorized API Key",
			requestType:  "GET",
			endpoint:     "/accounts/54400001111/transactions",
			headers:      map[string]string{"Authorization": "Bearer " + randomToken},
			expectedCode: http.StatusUnauthorized,
			expectedBody: expectedUnauthorizedBody,
		},
	}
	testReqests(t, setUpTestRouter(), getTests)

}

// TestAccountTransaction tests GET requests for the /accounts/:accountId/transactions/:transactionRef endpoint.
// This endpoint returns a specific account transaction.
func TestAccountTransaction(t *testing.T) {

	adminToken := auth.NewAPIKey(auth.ROLE_ADMIN, 1337).Token()
	accountToken := auth.NewAPIKey(auth.ROLE_ACCOUNT, 54400001111).Token()
	randomToken := auth.NewAPIKey(auth.ROLE_ACCOUNT, 1337).Token()

	expectedOKBody := map[string]string{
		"reference":            "",
		"amount":               "",
		"creditDebitIndicator": "",
		"status":               "",
		"bookingDate":          "",
		"valueDate":            "",
		"accountServicerRef":   "",
		"bankTransactionCode":  "",
		"entryDetails":         "",
	}

	expectedUnauthorizedBody := map[string]string{
		"error":   "Unauthorized",
		"message": "Your API key is not authorized to access the requested resource",
	}

	expectedNotFoundBody := map[string]string{
		"error":   "Not Found",
		"message": "transaction not found",
	}

	getTests := []TestRequest{
		{
			testName:     "Valid API Key (Admin)",
			requestType:  "GET",
			endpoint:     "/accounts/54400001111/transactions/JAMBO81518-0029248",
			headers:      map[string]string{"Authorization": "Bearer " + adminToken},
			expectedCode: http.StatusOK,
			expectedBody: expectedOKBody,
		},
		{
			testName:     "Valid API Key (Account Owner)",
			requestType:  "GET",
			endpoint:     "/accounts/54400001111/transactions/JAMBO81518-0029248",
			headers:      map[string]string{"Authorization": "Bearer " + accountToken},
			expectedCode: http.StatusOK,
			expectedBody: expectedOKBody,
		},
		{
			testName:     "Unauthorized API Key",
			requestType:  "GET",
			endpoint:     "/accounts/54400001111/transactions/JAMBO81518-0029248",
			headers:      map[string]string{"Authorization": "Bearer " + randomToken},
			expectedCode: http.StatusUnauthorized,
			expectedBody: expectedUnauthorizedBody,
		},
		{
			testName:     "Non existant transaction",
			requestType:  "GET",
			endpoint:     "/accounts/54400001111/transactions/NON-EXISTANT-TRANSACTION-1337",
			headers:      map[string]string{"Authorization": "Bearer " + accountToken},
			expectedCode: http.StatusNotFound,
			expectedBody: expectedNotFoundBody,
		},
		{
			testName:     "Non existant transaction (Unauthorized)",
			requestType:  "GET",
			endpoint:     "/accounts/54400001111/transactions/NON-EXISTANT-TRANSACTION-1337",
			headers:      map[string]string{"Authorization": "Bearer " + randomToken},
			expectedCode: http.StatusUnauthorized,
			expectedBody: expectedUnauthorizedBody,
		},
	}
	testReqests(t, setUpTestRouter(), getTests)

}
