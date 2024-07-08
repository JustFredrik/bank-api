package auth

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

const ROLE_ANY = "any"
const ROLE_ADMIN = "admin"
const ROLE_ACCOUNT = "account"
const EVENT_CREATE = "CREATE"
const EVENT_DELETE = "DELETE"
const EVENT_VALIDATE = "VALIDATE"
const AUTH_LOG_STRING = "[AUTH]"
const AUTH_DEBUG_LOG_STRING = "[AUTH-debug]"
const AUTH_DEBUG_WARNING_STRING = "[AUTH-debug] [WARNING]"

var keyTracker KeyPool = KeyPool{
	APIKeys: make(map[string]IAPIKey),
}

var validTokenCharacters = "abcdefghijkmnpqrstuvxyzABCDEFGHJKLMN123456789"
var validTokenCount = len(validTokenCharacters)

const tokenLength = 32

// Create random number generator
var rng *rand.Rand

func init() {

	// If in production randomize the seed
	if strings.ToLower(os.Getenv("AUTH_MODE")) != "debug" {
		rng = rand.New(rand.NewSource(time.Now().UnixNano()))
	} else {
		fmt.Printf("%s Running in \"debug\" mode. Switch to \"release\" mode in production", AUTH_DEBUG_WARNING_STRING)
		rng = rand.New(rand.NewSource(1337))

	}
}

func generateSecretToken() string {

	t := make([]byte, tokenLength)

	// Populate byte slice with random values from valid characters
	for i := range tokenLength {
		t[i] = validTokenCharacters[rng.Intn(validTokenCount)]
	}

	return string(t)

}

// Implemented to later be able to integrate pubsub / or other event pattern
func emit(event string, key IAPIKey) {
	log(event, key)
}

func log(event string, key IAPIKey) { // Used to log each Key handler

	fmt.Printf("%s %-20s | %-7s | %-7s |  accountId %-20d  |  token %-20s \n",
		AUTH_LOG_STRING,
		time.Now().Format(time.RFC822),
		event,
		key.Role(),
		key.AccountId(),
		key.Token(),
	)
}

func NewAPIKey(role string, accountId uint64) IAPIKey {

	// validate role string and accountId
	switch role {
	case ROLE_ADMIN:
		role = ROLE_ADMIN
	default:
		role = ROLE_ACCOUNT
	}
	if role == ROLE_ADMIN {
		accountId = 0
	}

	// Create APIKey
	key := BaseAPIKey{
		token:       generateSecretToken(),
		role:        role,
		accountId:   accountId,
		createdTime: time.Now().Unix(),
	}

	// If token colission in Tracker re-generate Token
	for {
		if _, ok := keyTracker.APIKeys[key.Token()]; !ok {
			break
		}
		key.token = generateSecretToken()
	}

	// Track API Key in Tracker struct
	keyTracker.APIKeys[key.Token()] = key
	keyTracker.TotalCount++

	// Emit event
	emit(EVENT_CREATE, key)

	return key
}

func KeyHasAccess(c *gin.Context, required_role string) bool {
	return true
}
