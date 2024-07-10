package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/justfredrik/bank-api/internal/auth"
	"github.com/magiconair/properties/assert"
)

func setUpTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	return SetUpRouter()
}

func TestGetPing(t *testing.T) {

	router := setUpTestRouter()

	// Create Auth token for test
	apiKey := auth.NewAPIKey(auth.ROLE_ACCOUNT, 112233445566778899)
	token := apiKey.Token()

	req, _ := http.NewRequest("GET", "/ping", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	print(req.Header.Get("Authorization"))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	var responseMap map[string]string
	json.Unmarshal(w.Body.Bytes(), &responseMap)

	assert.Equal(t, w.Code, http.StatusOK)
	assert.Equal(t, w.Body.String(), "{\"message\":\"pong\"}")
}
