package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/justfredrik/bank-api/internal/auth"
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

func testReqests(t *testing.T, router *gin.Engine, tests []TestRequest) {
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

			println(w.Body.String())

			// Validate Response Format
			bodyIsOK, err := jsonContains(w.Body.Bytes(), test.expectedBody)
			assert.Equal(t, true, bodyIsOK, err)

		})
	}
}

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

func TestAccounts(t *testing.T) {

	adminToken := auth.NewAPIKey(auth.ROLE_ADMIN, 1337).Token()
	accountToken := auth.NewAPIKey(auth.ROLE_ACCOUNT, 54400001111).Token()

	getTests := []TestRequest{
		{
			testName:     "	Valid API Key",
			requestType:  "GET",
			endpoint:     "/accounts",
			headers:      map[string]string{"Authorization": "Bearer " + adminToken},
			expectedCode: http.StatusOK,
			expectedBody: map[string]string{"accounts": ""},
		},
		{
			testName:     "	Unauthorized API Key",
			requestType:  "GET",
			endpoint:     "/accounts",
			headers:      map[string]string{"Authorization": "Bearer " + accountToken},
			expectedCode: http.StatusUnauthorized,
			expectedBody: map[string]string{"error": "Unauthorized", "message": "Your API key is not authorized to access the requested resource"},
		},
		{
			testName:     "	malformed header",
			requestType:  "GET",
			endpoint:     "/accounts",
			headers:      map[string]string{"Authorization": "Beer " + accountToken},
			expectedCode: http.StatusUnauthorized,
			expectedBody: map[string]string{"error": "Unauthorized", "message": "Malformed Authorization header"},
		},
		{
			testName:     "	missing header",
			requestType:  "GET",
			endpoint:     "/accounts",
			headers:      map[string]string{"ThisIsWrong": "Fox " + accountToken},
			expectedCode: http.StatusUnauthorized,
			expectedBody: map[string]string{"error": "Unauthorized", "message": "Missing Authorization header"},
		},
	}
	putTests := []TestRequest{}
	patchTests := []TestRequest{}
	deleteTests := []TestRequest{}

	testReqests(t, setUpTestRouter(), getTests)
	testReqests(t, setUpTestRouter(), putTests)
	testReqests(t, setUpTestRouter(), patchTests)
	testReqests(t, setUpTestRouter(), deleteTests)
}
